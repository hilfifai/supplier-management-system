<template>
  <div class="new-supplier">
    <div class="header">
      <h1>{{ isEdit ? 'Edit Supplier' : 'New Supplier' }}</h1>
      <div class="header-actions">
        <button @click="$parent.navigateTo('supplier-list')" class="btn btn-secondary">Back to List</button>
      </div>
    </div>

    <div class="form-container">
      <form @submit.prevent="saveSupplier">
        <div class="form-section">
          <h3>Basic Information</h3>
          <div class="form-row">
            <div class="form-group">
              <label>Supplier Name *</label>
              <input 
                v-model="supplier.name" 
                type="text" 
                required 
                class="form-control"
                placeholder="PT Setroom Indonesia"
              >
            </div>
            <div class="form-group">
              <label>Nick Name</label>
              <input 
                v-model="supplier.nick_name" 
                type="text" 
                class="form-control"
                placeholder="Setroom"
              >
            </div>
          </div>
          
          <div class="form-row">
            <div class="form-group">
              <label>Logo URL</label>
              <input 
                v-model="supplier.logo_url" 
                type="url" 
                class="form-control"
                placeholder="https://example.com/logo.png"
              >
            </div>
            <div class="form-group">
              <label>Status</label>
              <select v-model="supplier.status" class="form-control">
                <option value="Active">Active</option>
                <option value="In Progress">In Progress</option>
                <option value="Blocked">Blocked</option>
              </select>
            </div>
          </div>
        </div>

        <div class="form-section">
          <h3>Address</h3>
          <div v-for="(address, index) in supplier.addresses" :key="index" class="address-item">
            <div class="form-row">
              <div class="form-group">
                <label>Type</label>
                <select v-model="address.type" class="form-control">
                  <option value="Head Office">Head Office</option>
                  <option value="Branch Office">Branch Office</option>
                  <option value="Warehouse">Warehouse</option>
                </select>
              </div>
              <div class="form-group">
                <label>Street Address</label>
                <input 
                  v-model="address.street" 
                  type="text" 
                  class="form-control"
                  placeholder="Fatmawati Raya St, 123"
                >
              </div>
            </div>
            <div class="form-row">
              <div class="form-group">
                <label>City</label>
                <input 
                  v-model="address.city" 
                  type="text" 
                  class="form-control"
                  placeholder="Jakarta Selatan"
                >
              </div>
              <div class="form-group">
                <label>Province</label>
                <input 
                  v-model="address.province" 
                  type="text" 
                  class="form-control"
                  placeholder="DKI Jakarta"
                >
              </div>
            </div>
            <div class="form-row">
              <div class="form-group">
                <label>
                  <input 
                    v-model="address.is_main" 
                    type="checkbox"
                    @change="setMainAddress(index)"
                  >
                  Main Address
                </label>
              </div>
              <div class="form-group">
                <button 
                  type="button" 
                  @click="removeAddress(index)"
                  class="btn btn-danger btn-sm"
                  v-if="supplier.addresses.length > 1"
                >
                  Remove
                </button>
              </div>
            </div>
          </div>
          <button type="button" @click="addAddress" class="btn btn-secondary">Add Address</button>
        </div>

        <div class="form-section">
          <h3>Contacts</h3>
          <div v-for="(contact, index) in supplier.contacts" :key="index" class="contact-item">
            <div class="form-row">
              <div class="form-group">
                <label>Name *</label>
                <input 
                  v-model="contact.name" 
                  type="text" 
                  required 
                  class="form-control"
                  placeholder="Albert Einstein"
                >
              </div>
              <div class="form-group">
                <label>Job Position</label>
                <input 
                  v-model="contact.job_position" 
                  type="text" 
                  class="form-control"
                  placeholder="CEO"
                >
              </div>
            </div>
            <div class="form-row">
              <div class="form-group">
                <label>Email</label>
                <input 
                  v-model="contact.email" 
                  type="email" 
                  class="form-control"
                  placeholder="einstein@gmail.com"
                >
              </div>
              <div class="form-group">
                <label>Phone</label>
                <input 
                  v-model="contact.phone" 
                  type="text" 
                  class="form-control"
                  placeholder="021.123456"
                >
              </div>
            </div>
            <div class="form-row">
              <div class="form-group">
                <label>Mobile</label>
                <input 
                  v-model="contact.mobile" 
                  type="text" 
                  class="form-control"
                  placeholder="0811234567"
                >
              </div>
              <div class="form-group">
                <label>
                  <input 
                    v-model="contact.is_main" 
                    type="checkbox"
                    @change="setMainContact(index)"
                  >
                  Main Contact
                </label>
              </div>
            </div>
            <div class="form-row">
              <div class="form-group">
                <button 
                  type="button" 
                  @click="removeContact(index)"
                  class="btn btn-danger btn-sm"
                  v-if="supplier.contacts.length > 1"
                >
                  Remove
                </button>
              </div>
            </div>
          </div>
          <button type="button" @click="addContact" class="btn btn-secondary">Add Contact</button>
        </div>

        <div class="form-actions">
          <button type="button" @click="$parent.navigateTo('supplier-list')" class="btn btn-secondary">Cancel</button>
          <button type="submit" class="btn btn-primary">{{ isEdit ? 'Update' : 'Save' }}</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'NewSupplier',
  props: ['supplierData'],
  data() {
    return {
      supplier: {
        name: '',
        nick_name: '',
        logo_url: '',
        status: 'Active',
        addresses: [
          {
            type: 'Head Office',
            street: '',
            city: '',
            province: '',
            postal_code: '',
            is_main: true
          }
        ],
        contacts: [
          {
            name: '',
            job_position: '',
            email: '',
            phone: '',
            mobile: '',
            is_main: true
          }
        ]
      },
      isEdit: false
    }
  },
  mounted() {
    if (this.$parent.selectedSupplier) {
      this.supplier = { ...this.$parent.selectedSupplier };
      this.isEdit = true;
      
      if (!this.supplier.addresses || this.supplier.addresses.length === 0) {
        this.supplier.addresses = [
          {
            type: 'Head Office',
            street: '',
            city: '',
            province: '',
            postal_code: '',
            is_main: true
          }
        ];
      }
      
      if (!this.supplier.contacts || this.supplier.contacts.length === 0) {
        this.supplier.contacts = [
          {
            name: '',
            job_position: '',
            email: '',
            phone: '',
            mobile: '',
            is_main: true
          }
        ];
      }
    }
  },
  methods: {
    addAddress() {
      this.supplier.addresses.push({
        type: 'Branch Office',
        street: '',
        city: '',
        province: '',
        postal_code: '',
        is_main: false
      });
    },
    removeAddress(index) {
      this.supplier.addresses.splice(index, 1);
    },
    setMainAddress(index) {
      this.supplier.addresses.forEach((addr, i) => {
        addr.is_main = i === index;
      });
    },
    addContact() {
      this.supplier.contacts.push({
        name: '',
        job_position: '',
        email: '',
        phone: '',
        mobile: '',
        is_main: false
      });
    },
    removeContact(index) {
      this.supplier.contacts.splice(index, 1);
    },
    setMainContact(index) {
      this.supplier.contacts.forEach((contact, i) => {
        contact.is_main = i === index;
      });
    },
    async saveSupplier() {
      try {
        if (this.isEdit) {
          await axios.put(`/suppliers/${this.supplier.id}`, this.supplier);
        } else {
          await axios.post('/suppliers', this.supplier);
        }
        
        this.$parent.navigateTo('supplier-list');
      } catch (error) {
        console.error('Error saving supplier:', error);
        alert('Error saving supplier. Please try again.');
      }
    }
  }
}
</script>

<style scoped>
.new-supplier {
  padding: 20px;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
}

.header h1 {
  color: #333;
  margin: 0;
}

.form-container {
  background: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.form-section {
  margin-bottom: 30px;
  padding-bottom: 20px;
  border-bottom: 1px solid #eee;
}

.form-section:last-child {
  border-bottom: none;
}

.form-section h3 {
  color: #333;
  margin-bottom: 20px;
  font-size: 18px;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
  margin-bottom: 15px;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-group label {
  margin-bottom: 5px;
  font-weight: 500;
  color: #333;
}

.form-control {
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.form-control:focus {
  outline: none;
  border-color: #007bff;
  box-shadow: 0 0 0 2px rgba(0,123,255,0.25);
}

.address-item, .contact-item {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 6px;
  margin-bottom: 15px;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 15px;
  margin-top: 30px;
  padding-top: 20px;
  border-top: 1px solid #eee;
}

.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
}

.btn-sm {
  padding: 6px 12px;
  font-size: 12px;
}

.btn-primary { background-color: #007bff; color: white; }
.btn-secondary { background-color: #6c757d; color: white; }
.btn-danger { background-color: #dc3545; color: white; }

.btn:hover {
  opacity: 0.9;
}

input[type="checkbox"] {
  margin-right: 8px;
}
</style>


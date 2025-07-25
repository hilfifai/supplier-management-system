<template>
  <div class="supplier-detail" v-if="supplier">
    <div class="header">
      <div class="supplier-header">
        <div class="supplier-logo">
          <img v-if="supplier.logo_url" :src="supplier.logo_url" :alt="supplier.name">
          <div v-else class="logo-placeholder">{{ supplier.nick_name?.charAt(0) || 'S' }}</div>
        </div>
        <div class="supplier-info">
          <h1>{{ supplier.name }}</h1>
          <div class="supplier-meta">
            <span class="supplier-id">{{ supplier.id }}</span>
            <span :class="['status-badge', getStatusClass(supplier.status)]">
              {{ supplier.status }}
            </span>
          </div>
          <div class="supplier-address">{{ getMainAddress() }}</div>
        </div>
      </div>
      <div class="header-actions">
        <button @click="$parent.navigateTo('supplier-list')" class="btn btn-secondary">Back to List</button>
        <button @click="editSupplier" class="btn btn-warning">Edit</button>
        <button 
          v-if="supplier.status !== 'Blocked'"
          @click="blockSupplier" 
          class="btn btn-danger"
        >
          Block
        </button>
        <button 
          v-else
          @click="unblockSupplier" 
          class="btn btn-success"
        >
          Unblock
        </button>
      </div>
    </div>

    <div class="tabs">
      <button 
        v-for="tab in tabs" 
        :key="tab.id"
        @click="activeTab = tab.id"
        :class="['tab-button', { active: activeTab === tab.id }]"
      >
        {{ tab.label }}
      </button>
    </div>

    <div class="tab-content">
      <!-- Overview Tab -->
      <div v-if="activeTab === 'overview'" class="tab-pane">
        <div class="overview-grid">
          <div class="info-section">
            <h3>Addresses</h3>
            <div v-for="address in supplier.addresses" :key="address.id" class="info-item">
              <div class="info-label">{{ address.type }}</div>
              <div class="info-value">
                {{ address.street }}<br>
                {{ address.city }}, {{ address.province }}
                <span v-if="address.is_main" class="main-badge">Main</span>
              </div>
            </div>
          </div>

          <div class="info-section">
            <h3>Contacts</h3>
            <div v-for="contact in supplier.contacts" :key="contact.id" class="info-item">
              <div class="info-label">{{ contact.name }}</div>
              <div class="info-value">
                {{ contact.job_position }}<br>
                {{ contact.email }}<br>
                {{ contact.phone }} / {{ contact.mobile }}
                <span v-if="contact.is_main" class="main-badge">Main</span>
              </div>
            </div>
          </div>

          <div class="info-section">
            <h3>Groups</h3>
            <div v-for="group in supplier.groups" :key="group.id" class="info-item">
              <div class="info-label">{{ group.group_name }}</div>
              <div class="info-value">{{ group.value }}</div>
            </div>
          </div>

          <div class="info-section">
            <h3>Other Information</h3>
            <div v-for="info in supplier.other_informations" :key="info.id" class="info-item">
              <div class="info-label">{{ info.attribute_name }}</div>
              <div class="info-value">{{ info.value }}</div>
            </div>
          </div>
        </div>
      </div>

      <!-- Materials Tab -->
      <div v-if="activeTab === 'materials'" class="tab-pane">
        <div class="materials-header">
          <h3>Materials Catalog</h3>
          <button class="btn btn-primary">Add Material</button>
        </div>
        <div class="table-container">
          <table class="materials-table">
            <thead>
              <tr>
                <th>Material ID</th>
                <th>Material Name</th>
                <th>Group</th>
                <th>Price</th>
                <th>Status</th>
                <th>Actions</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="material in supplier.materials" :key="material.id">
                <td>{{ material.material_id }}</td>
                <td>{{ material.material_name || 'N/A' }}</td>
                <td>{{ material.material_group }}</td>
                <td>Rp {{ formatCurrency(material.price) }}</td>
                <td>
                  <span :class="['status-badge', material.is_active ? 'status-active' : 'status-inactive']">
                    {{ material.is_active ? 'Active' : 'Inactive' }}
                  </span>
                </td>
                <td>
                  <button class="btn btn-sm btn-primary">Edit</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Orders Tab -->
      <div v-if="activeTab === 'orders'" class="tab-pane">
        <div class="orders-stats">
          <div class="stat-card">
            <h4>Total Orders</h4>
            <div class="stat-value">{{ orderStats.total || 4257 }}</div>
          </div>
          <div class="stat-card">
            <h4>On Time Delivery</h4>
            <div class="stat-value">{{ orderStats.onTime || 1869 }}</div>
          </div>
          <div class="stat-card">
            <h4>Late Delivery</h4>
            <div class="stat-value">{{ orderStats.late || 999 }}</div>
          </div>
          <div class="stat-card">
            <h4>In Progress</h4>
            <div class="stat-value">{{ orderStats.inProgress || 999 }}</div>
          </div>
        </div>
        
        <div class="table-container">
          <table class="orders-table">
            <thead>
              <tr>
                <th>Order ID (PO)</th>
                <th>Shipment Date</th>
                <th>Order Status</th>
                <th>Estimated Delivery Date</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="order in supplier.order_histories" :key="order.id">
                <td>{{ order.order_id_po }}</td>
                <td>{{ formatDate(order.shipment_date) }}</td>
                <td>{{ order.order_status }}</td>
                <td>{{ formatDate(order.estimated_delivery_date) }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Invoices Tab -->
      <div v-if="activeTab === 'invoices'" class="tab-pane">
        <div class="invoices-stats">
          <div class="stat-card">
            <h4>Total Invoices</h4>
            <div class="stat-value">{{ invoiceStats.total || 4257 }}</div>
          </div>
          <div class="stat-card">
            <h4>In Progress</h4>
            <div class="stat-value">{{ invoiceStats.inProgress || 1869 }}</div>
          </div>
          <div class="stat-card">
            <h4>Paid</h4>
            <div class="stat-value">{{ invoiceStats.paid || 999 }}</div>
          </div>
          <div class="stat-card">
            <h4>Outstanding</h4>
            <div class="stat-value">{{ invoiceStats.outstanding || 999 }}</div>
          </div>
        </div>
        
        <div class="table-container">
          <table class="invoices-table">
            <thead>
              <tr>
                <th>Invoice No</th>
                <th>Project Name</th>
                <th>Amount</th>
                <th>Aging (days)</th>
                <th>Received Date</th>
                <th>Est. Payment Date</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="invoice in supplier.invoices" :key="invoice.id">
                <td>{{ invoice.invoice_number }}</td>
                <td>{{ invoice.project_name }}</td>
                <td>Rp {{ formatCurrency(invoice.amount) }}</td>
                <td>{{ invoice.aging_days }}</td>
                <td>{{ formatDate(invoice.received_date) }}</td>
                <td>{{ formatDate(invoice.estimated_payment_date) }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'SupplierDetail',
  data() {
    return {
      supplier: null,
      activeTab: 'overview',
      tabs: [
        { id: 'overview', label: 'Overview' },
        { id: 'materials', label: 'Material Catalog' },
        { id: 'orders', label: 'Orders' },
        { id: 'invoices', label: 'Invoices' }
      ],
      orderStats: {},
      invoiceStats: {}
    }
  },
  mounted() {
    if (this.$parent.selectedSupplier) {
      this.loadSupplierDetail(this.$parent.selectedSupplier.id);
    }
  },
  methods: {
    async loadSupplierDetail(id) {
      try {
        const response = await axios.get(`/suppliers/${id}`);
        this.supplier = response.data.data;
      } catch (error) {
        console.error('Error loading supplier detail:', error);
        this.supplier = this.getMockSupplierDetail();
      }
    },
    getMockSupplierDetail() {
      return {
        id: 'STRM61000012',
        name: 'PT Setroom Indonesia',
        nick_name: 'Setroom',
        status: 'Active',
        logo_url: '',
        addresses: [
          {
            id: 'addr1',
            type: 'Head Office',
            street: 'Fatmawati Raya St, 123',
            city: 'Jakarta Selatan',
            province: 'DKI Jakarta',
            is_main: true
          }
        ],
        contacts: [
          {
            id: 'contact1',
            name: 'Albert Einstein',
            job_position: 'CEO',
            email: 'einstein@gmail.com',
            phone: '021.123456',
            mobile: '0811234567',
            is_main: true
          }
        ],
        groups: [
          {
            id: 'group1',
            group_name: 'Industry',
            value: 'Manufacture'
          }
        ],
        materials: [
          {
            id: 'mat1',
            material_id: '83E30041UK',
            material_name: 'Lenovo Yoga #83E30041UK',
            material_group: 'Computer / Notebook',
            price: 19000000,
            is_active: true
          }
        ],
        other_informations: [
          {
            id: 'other1',
            attribute_name: 'SAP Vendor Id',
            value: '4100012'
          },
          {
            id: 'other2',
            attribute_name: 'NPWP',
            value: '1.000.000-123'
          }
        ],
        order_histories: [],
        invoices: []
      };
    },
    editSupplier() {
      this.$parent.navigateTo('new-supplier', this.supplier);
    },
    async blockSupplier() {
      const reason = prompt('Enter reason for blocking:');
      if (reason) {
        try {
          await axios.post(`/suppliers/${this.supplier.id}/block`, { reason });
          this.supplier.status = 'Blocked';
        } catch (error) {
          console.error('Error blocking supplier:', error);
        }
      }
    },
    async unblockSupplier() {
      const reason = prompt('Enter reason for unblocking:');
      if (reason) {
        try {
          await axios.post(`/suppliers/${this.supplier.id}/unblock`, { reason });
          this.supplier.status = 'Active';
        } catch (error) {
          console.error('Error unblocking supplier:', error);
        }
      }
    },
    getMainAddress() {
      if (this.supplier.addresses && this.supplier.addresses.length > 0) {
        const mainAddress = this.supplier.addresses.find(addr => addr.is_main) || this.supplier.addresses[0];
        return `${mainAddress.street}, ${mainAddress.city}`;
      }
      return 'N/A';
    },
    getStatusClass(status) {
      switch (status) {
        case 'Active': return 'status-active';
        case 'In Progress': return 'status-progress';
        case 'Blocked': return 'status-blocked';
        default: return '';
      }
    },
    formatCurrency(value) {
      return new Intl.NumberFormat('id-ID').format(value);
    },
    formatDate(date) {
      if (!date) return 'N/A';
      return new Date(date).toLocaleDateString('id-ID');
    }
  }
}
</script>

<style scoped>
.supplier-detail {
  padding: 20px;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 30px;
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.supplier-header {
  display: flex;
  gap: 20px;
}

.supplier-logo {
  width: 80px;
  height: 80px;
  border-radius: 8px;
  overflow: hidden;
  background: #f8f9fa;
  display: flex;
  align-items: center;
  justify-content: center;
}

.supplier-logo img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.logo-placeholder {
  font-size: 32px;
  font-weight: bold;
  color: #666;
}

.supplier-info h1 {
  margin: 0 0 10px 0;
  color: #333;
}

.supplier-meta {
  display: flex;
  gap: 15px;
  align-items: center;
  margin-bottom: 10px;
}

.supplier-id {
  font-family: monospace;
  background: #f8f9fa;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
}

.supplier-address {
  color: #666;
  font-size: 14px;
}

.header-actions {
  display: flex;
  gap: 10px;
}

.tabs {
  display: flex;
  background: white;
  border-radius: 8px 8px 0 0;
  overflow: hidden;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.tab-button {
  padding: 15px 25px;
  border: none;
  background: #f8f9fa;
  cursor: pointer;
  font-weight: 500;
  color: #666;
  border-bottom: 3px solid transparent;
}

.tab-button.active {
  background: white;
  color: #007bff;
  border-bottom-color: #007bff;
}

.tab-content {
  background: white;
  padding: 30px;
  border-radius: 0 0 8px 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.overview-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 30px;
}

.info-section h3 {
  color: #333;
  margin-bottom: 20px;
  font-size: 18px;
}

.info-item {
  margin-bottom: 15px;
  padding-bottom: 15px;
  border-bottom: 1px solid #eee;
}

.info-item:last-child {
  border-bottom: none;
}

.info-label {
  font-weight: 500;
  color: #333;
  margin-bottom: 5px;
}

.info-value {
  color: #666;
  line-height: 1.5;
}

.main-badge {
  background: #007bff;
  color: white;
  padding: 2px 6px;
  border-radius: 10px;
  font-size: 10px;
  margin-left: 10px;
}

.materials-header, .orders-header, .invoices-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.orders-stats, .invoices-stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
  margin-bottom: 30px;
}

.stat-card {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  text-align: center;
}

.stat-card h4 {
  margin: 0 0 10px 0;
  color: #666;
  font-size: 14px;
}

.stat-value {
  font-size: 24px;
  font-weight: bold;
  color: #333;
}

.table-container {
  background: white;
  border-radius: 8px;
  overflow: hidden;
  border: 1px solid #dee2e6;
}

.materials-table, .orders-table, .invoices-table {
  width: 100%;
  border-collapse: collapse;
}

.materials-table th, .orders-table th, .invoices-table th {
  background-color: #f8f9fa;
  padding: 12px;
  text-align: left;
  font-weight: 600;
  color: #333;
  border-bottom: 1px solid #dee2e6;
}

.materials-table td, .orders-table td, .invoices-table td {
  padding: 12px;
  border-bottom: 1px solid #dee2e6;
}

.status-badge {
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
}

.status-active {
  background-color: #d4edda;
  color: #155724;
}

.status-progress {
  background-color: #fff3cd;
  color: #856404;
}

.status-blocked {
  background-color: #f8d7da;
  color: #721c24;
}

.status-inactive {
  background-color: #f8f9fa;
  color: #6c757d;
}

.btn {
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
}

.btn-sm {
  padding: 4px 8px;
  font-size: 12px;
}

.btn-primary { background-color: #007bff; color: white; }
.btn-secondary { background-color: #6c757d; color: white; }
.btn-success { background-color: #28a745; color: white; }
.btn-warning { background-color: #ffc107; color: #212529; }
.btn-danger { background-color: #dc3545; color: white; }
</style>


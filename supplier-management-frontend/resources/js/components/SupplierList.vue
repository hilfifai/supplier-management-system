<template>
  <div class="supplier-list">
    <div class="header">
      <h1>Supplier List</h1>
      <div class="header-actions">
        <button @click="$parent.navigateTo('dashboard')" class="btn btn-secondary">Dashboard</button>
        <button @click="$parent.navigateTo('new-supplier')" class="btn btn-success">New Supplier</button>
        <button @click="exportData" class="btn btn-info">Export</button>
      </div>
    </div>

    <div class="filters">
      <input 
        v-model="searchQuery" 
        @input="searchSuppliers"
        type="text" 
        placeholder="Search Customer" 
        class="search-input"
      >
      <select v-model="statusFilter" @change="filterSuppliers" class="status-filter">
        <option value="">All Status</option>
        <option value="Active">Active</option>
        <option value="In Progress">In Progress</option>
        <option value="Blocked">Blocked</option>
      </select>
    </div>

    <div class="table-container">
      <table class="suppliers-table">
        <thead>
          <tr>
            <th>#</th>
            <th>Name</th>
            <th>Address</th>
            <th>Contact</th>
            <th>Status</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(supplier, index) in suppliers" :key="supplier.id">
            <td>{{ index + 1 }}</td>
            <td>
              <div class="supplier-info">
                <div class="supplier-name">{{ supplier.name }}</div>
                <div class="supplier-code">{{ supplier.id }} [{{ supplier.nick_name }}]</div>
              </div>
            </td>
            <td>{{ getMainAddress(supplier) }}</td>
            <td>{{ getMainContact(supplier) }}</td>
            <td>
              <span :class="['status-badge', getStatusClass(supplier.status)]">
                {{ supplier.status }}
              </span>
            </td>
            <td>
              <div class="actions">
                <button @click="viewSupplier(supplier)" class="btn btn-sm btn-primary">View</button>
                <button @click="editSupplier(supplier)" class="btn btn-sm btn-warning">Edit</button>
                <button 
                  v-if="supplier.status !== 'Blocked'"
                  @click="blockSupplier(supplier)" 
                  class="btn btn-sm btn-danger"
                >
                  Block
                </button>
                <button 
                  v-else
                  @click="unblockSupplier(supplier)" 
                  class="btn btn-sm btn-success"
                >
                  Unblock
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div class="pagination" v-if="pagination.total_pages > 1">
      <button 
        @click="changePage(pagination.page - 1)"
        :disabled="pagination.page <= 1"
        class="btn btn-sm"
      >
        Previous
      </button>
      <span class="page-info">
        Page {{ pagination.page }} of {{ pagination.total_pages }}
      </span>
      <button 
        @click="changePage(pagination.page + 1)"
        :disabled="pagination.page >= pagination.total_pages"
        class="btn btn-sm"
      >
        Next
      </button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'SupplierList',
  data() {
    return {
      suppliers: [],
      searchQuery: '',
      statusFilter: '',
      pagination: {
        page: 1,
        limit: 10,
        total: 0,
        total_pages: 0
      }
    }
  },
  mounted() {
    this.loadSuppliers();
  },
  methods: {
    async loadSuppliers() {
      try {
        const params = {
          page: this.pagination.page,
          limit: this.pagination.limit
        };
        
        if (this.searchQuery) {
          params.search = this.searchQuery;
        }
        
        if (this.statusFilter) {
          params.status = this.statusFilter;
        }

        const response = await axios.get('/suppliers', { params });
        this.suppliers = response.data.data;
        this.pagination = response.data.pagination;
      } catch (error) {
        console.error('Error loading suppliers:', error);
        this.suppliers = this.getMockSuppliers();
      }
    },
    getMockSuppliers() {
      return [
        {
          id: 'STRM61000012',
          name: 'PT Setroom Indonesia',
          nick_name: 'Setroom',
          status: 'Active',
          addresses: [{ street: 'Jakarta, Indonesia', is_main: true }],
          contacts: [{ name: 'Albert Einstein', is_main: true }]
        },
        {
          id: 'SKSK41000013',
          name: 'PT Suka Suka',
          nick_name: 'Sukasuka',
          status: 'In Progress',
          addresses: [{ street: 'Bandung, Indonesia', is_main: true }],
          contacts: [{ name: 'James Lee', is_main: true }]
        },
        {
          id: 'ARIB41000014',
          name: 'PT Angin Ribut',
          nick_name: 'Angin',
          status: 'Blocked',
          addresses: [{ street: 'Denpasar, Indonesia', is_main: true }],
          contacts: [{ name: 'Maria Chen', is_main: true }]
        }
      ];
    },
    searchSuppliers() {
      this.pagination.page = 1;
      this.loadSuppliers();
    },
    filterSuppliers() {
      this.pagination.page = 1;
      this.loadSuppliers();
    },
    changePage(page) {
      this.pagination.page = page;
      this.loadSuppliers();
    },
    viewSupplier(supplier) {
      this.$parent.navigateTo('supplier-detail', supplier);
    },
    editSupplier(supplier) {
      this.$parent.navigateTo('new-supplier', supplier);
    },
    async blockSupplier(supplier) {
      const reason = prompt('Enter reason for blocking:');
      if (reason) {
        try {
          await axios.post(`/suppliers/${supplier.id}/block`, { reason });
          this.loadSuppliers();
        } catch (error) {
          console.error('Error blocking supplier:', error);
        }
      }
    },
    async unblockSupplier(supplier) {
      const reason = prompt('Enter reason for unblocking:');
      if (reason) {
        try {
          await axios.post(`/suppliers/${supplier.id}/unblock`, { reason });
          this.loadSuppliers();
        } catch (error) {
          console.error('Error unblocking supplier:', error);
        }
      }
    },
    exportData() {
      console.log('Export functionality would be implemented here');
    },
    getMainAddress(supplier) {
      if (supplier.addresses && supplier.addresses.length > 0) {
        const mainAddress = supplier.addresses.find(addr => addr.is_main) || supplier.addresses[0];
        return mainAddress.street || mainAddress.city || 'N/A';
      }
      return 'N/A';
    },
    getMainContact(supplier) {
      if (supplier.contacts && supplier.contacts.length > 0) {
        const mainContact = supplier.contacts.find(contact => contact.is_main) || supplier.contacts[0];
        return mainContact.name || 'N/A';
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
    }
  }
}
</script>

<style scoped>
.supplier-list {
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

.header-actions {
  display: flex;
  gap: 10px;
}

.filters {
  display: flex;
  gap: 15px;
  margin-bottom: 20px;
}

.search-input, .status-filter {
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.search-input {
  flex: 1;
  max-width: 300px;
}

.table-container {
  background: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.suppliers-table {
  width: 100%;
  border-collapse: collapse;
}

.suppliers-table th {
  background-color: #f8f9fa;
  padding: 12px;
  text-align: left;
  font-weight: 600;
  color: #333;
  border-bottom: 1px solid #dee2e6;
}

.suppliers-table td {
  padding: 12px;
  border-bottom: 1px solid #dee2e6;
}

.supplier-info {
  display: flex;
  flex-direction: column;
}

.supplier-name {
  font-weight: 500;
  color: #333;
}

.supplier-code {
  font-size: 12px;
  color: #666;
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

.actions {
  display: flex;
  gap: 5px;
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
.btn-info { background-color: #17a2b8; color: white; }

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 15px;
  margin-top: 20px;
}

.page-info {
  color: #666;
  font-size: 14px;
}
</style>


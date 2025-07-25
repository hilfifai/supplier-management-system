<template>
  <div class="dashboard">
    <div class="header">
      <h1>Supplier Management</h1>
      <div class="header-actions">
        <button @click="$parent.navigateTo('supplier-list')" class="btn btn-primary">Supplier List</button>
        <button @click="$parent.navigateTo('new-supplier')" class="btn btn-success">New Supplier</button>
      </div>
    </div>
    
    <div class="stats-grid">
      <div class="stat-card">
        <h3>Total Supplier</h3>
        <div class="stat-value">{{ stats.total_suppliers || 1869 }}</div>
        <div class="stat-change positive">+8% vs last year</div>
      </div>
      
      <div class="stat-card">
        <h3>New Supplier</h3>
        <div class="stat-value">{{ stats.new_suppliers || 27 }}</div>
        <div class="stat-change positive">+1% vs Last Year</div>
      </div>
      
      <div class="stat-card">
        <h3>Avg Cost per Supplier</h3>
        <div class="stat-value">Rp {{ formatCurrency(stats.avg_cost_per_supplier || 320300000) }}</div>
        <div class="stat-change negative">-1% vs Last Year</div>
      </div>
      
      <div class="stat-card">
        <h3>Blocked Supplier</h3>
        <div class="stat-value">{{ stats.blocked_suppliers || 31 }}</div>
        <div class="stat-change negative">-4% vs Last Year</div>
      </div>
    </div>

    <div class="orders-stats">
      <h2>Order Statistics</h2>
      <div class="stats-grid">
        <div class="stat-card">
          <h3>Total Order</h3>
          <div class="stat-value">{{ stats.total_orders || 4257 }}</div>
        </div>
        
        <div class="stat-card">
          <h3>On Time Delivery</h3>
          <div class="stat-value">{{ stats.on_time_delivery || 1869 }}</div>
        </div>
        
        <div class="stat-card">
          <h3>Late Delivery</h3>
          <div class="stat-value">{{ stats.late_delivery || 999 }}</div>
        </div>
        
        <div class="stat-card">
          <h3>Order In Progress</h3>
          <div class="stat-value">{{ stats.orders_in_progress || 999 }}</div>
        </div>
      </div>
    </div>

    <div class="invoice-stats">
      <h2>Invoice Statistics</h2>
      <div class="stats-grid">
        <div class="stat-card">
          <h3>Total Invoices</h3>
          <div class="stat-value">{{ stats.total_invoices || 4257 }}</div>
        </div>
        
        <div class="stat-card">
          <h3>In Progress</h3>
          <div class="stat-value">{{ stats.invoices_in_progress || 1869 }}</div>
        </div>
        
        <div class="stat-card">
          <h3>Paid</h3>
          <div class="stat-value">{{ stats.paid_invoices || 999 }}</div>
        </div>
        
        <div class="stat-card">
          <h3>Outstanding</h3>
          <div class="stat-value">{{ stats.outstanding_invoices || 999 }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Dashboard',
  data() {
    return {
      stats: {}
    }
  },
  mounted() {
    this.loadStats();
  },
  methods: {
     navigateTo(page) {
      this.$emit('navigate', page);
    },
    async loadStats() {
      try {
        const response = await axios.get('/dashboard/stats');
        this.stats = response.data.data;
      } catch (error) {
        console.error('Error loading stats:', error);
      }
    },
    formatCurrency(value) {
      return new Intl.NumberFormat('id-ID').format(value / 1000000) + ' Mn';
    }
  }
}
</script>

<style scoped>
.dashboard {
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

.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-weight: 500;
}

.btn-primary {
  background-color: #007bff;
  color: white;
}

.btn-success {
  background-color: #28a745;
  color: white;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
  margin-bottom: 30px;
}

.stat-card {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  border-left: 4px solid #007bff;
}

.stat-card h3 {
  margin: 0 0 10px 0;
  color: #666;
  font-size: 14px;
  font-weight: 500;
}

.stat-value {
  font-size: 24px;
  font-weight: bold;
  color: #333;
  margin-bottom: 5px;
}

.stat-change {
  font-size: 12px;
  font-weight: 500;
}

.stat-change.positive {
  color: #28a745;
}

.stat-change.negative {
  color: #dc3545;
}

.orders-stats, .invoice-stats {
  margin-bottom: 30px;
}

.orders-stats h2, .invoice-stats h2 {
  color: #333;
  margin-bottom: 20px;
}
</style>


<template>
  <div class="review-approvals">
    <div class="header">
      <h1>Review & Approvals</h1>
      <div class="header-actions">
        <button @click="$parent.navigateTo('dashboard')" class="btn btn-secondary">Dashboard</button>
      </div>
    </div>

    <div class="process-tabs">
      <button 
        v-for="process in processes" 
        :key="process.id"
        @click="activeProcess = process.id"
        :class="['process-tab', { active: activeProcess === process.id }]"
      >
        {{ process.label }}
      </button>
    </div>

    <div class="filters">
      <input 
        v-model="searchQuery" 
        @input="searchProcesses"
        type="text" 
        placeholder="Search Supplier" 
        class="search-input"
      >
    </div>

    <div class="table-container">
      <table class="processes-table">
        <thead>
          <tr>
            <th>Process Name</th>
            <th>Customer ID</th>
            <th>Customer Name</th>
            <th>Stage / Flow</th>
            <th>Status</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="process in filteredProcesses" :key="process.id">
            <td>{{ process.process_name }}</td>
            <td>{{ process.customer_id }}</td>
            <td>{{ process.customer_name }}</td>
            <td>{{ process.stage_flow }}</td>
            <td>
              <span :class="['status-badge', getStatusClass(process.status)]">
                {{ process.status }}
              </span>
            </td>
            <td>
              <div class="actions">
                <button @click="viewProcess(process)" class="btn btn-sm btn-primary">View</button>
                <button @click="approveProcess(process)" class="btn btn-sm btn-success">Approve</button>
                <button @click="rejectProcess(process)" class="btn btn-sm btn-danger">Reject</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Process Detail Modal -->
    <div v-if="selectedProcess" class="modal-overlay" @click="closeModal">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h2>{{ selectedProcess.process_name }}</h2>
          <button @click="closeModal" class="close-btn">&times;</button>
        </div>
        
        <div class="modal-body">
          <div class="process-info">
            <div class="info-row">
              <label>Customer:</label>
              <span>{{ selectedProcess.customer_name }} ({{ selectedProcess.customer_id }})</span>
            </div>
            <div class="info-row">
              <label>Stage:</label>
              <span>{{ selectedProcess.stage_flow }}</span>
            </div>
            <div class="info-row">
              <label>Status:</label>
              <span :class="['status-badge', getStatusClass(selectedProcess.status)]">
                {{ selectedProcess.status }}
              </span>
            </div>
          </div>

          <div class="workflow-section" v-if="selectedProcess.workflows">
            <h3>Workflow Steps</h3>
            <div class="workflow-steps">
              <div 
                v-for="workflow in selectedProcess.workflows" 
                :key="workflow.id"
                class="workflow-step"
              >
                <div class="step-name">{{ workflow.stage_name }}</div>
                <div class="step-sla">SLA: {{ workflow.sla_value }} {{ workflow.sla_unit }}</div>
                <div class="step-status">
                  <span :class="['status-badge', workflow.is_active ? 'status-active' : 'status-inactive']">
                    {{ workflow.is_active ? 'Active' : 'Completed' }}
                  </span>
                </div>
              </div>
            </div>
          </div>

          <div class="approval-section">
            <h3>Approval Action</h3>
            <div class="approval-form">
              <textarea 
                v-model="approvalNotes" 
                placeholder="Enter approval notes..."
                class="approval-notes"
              ></textarea>
              <div class="approval-actions">
                <button @click="submitApproval('approved')" class="btn btn-success">Approve</button>
                <button @click="submitApproval('rejected')" class="btn btn-danger">Reject</button>
                <button @click="submitApproval('returned')" class="btn btn-warning">Return</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'ReviewApprovals',
  data() {
    return {
      activeProcess: 'supplier-creation',
      processes: [
        { id: 'supplier-creation', label: 'Supplier Creation' },
        { id: 'supplier-assessment', label: 'Supplier Assessment' },
        { id: 'block-unblock', label: 'Block / Unblock Supplier' }
      ],
      searchQuery: '',
      reviewProcesses: [],
      selectedProcess: null,
      approvalNotes: ''
    }
  },
  computed: {
    filteredProcesses() {
      let filtered = this.reviewProcesses.filter(process => 
        process.process_name.toLowerCase().includes(this.activeProcess.replace('-', ' '))
      );
      
      if (this.searchQuery) {
        filtered = filtered.filter(process =>
          process.customer_name.toLowerCase().includes(this.searchQuery.toLowerCase()) ||
          process.customer_id.toLowerCase().includes(this.searchQuery.toLowerCase())
        );
      }
      
      return filtered;
    }
  },
  mounted() {
    this.loadReviewProcesses();
  },
  methods: {
    async loadReviewProcesses() {
      try {
        const response = await axios.get('/review-approvals');
        this.reviewProcesses = response.data.data;
      } catch (error) {
        console.error('Error loading review processes:', error);
        this.reviewProcesses = this.getMockProcesses();
      }
    },
    getMockProcesses() {
      return [
        {
          id: 'rev1',
          process_name: 'Supplier Creation',
          customer_id: '1001',
          customer_name: 'PT Suka Suka',
          stage_flow: '01-02-03-05',
          status: '02: Procurement',
          workflows: [
            {
              id: 'wf1',
              stage_name: 'Draft',
              sla_value: '24',
              sla_unit: 'hours',
              is_active: false
            },
            {
              id: 'wf2',
              stage_name: 'In Review',
              sla_value: '24',
              sla_unit: 'hours',
              is_active: true
            },
            {
              id: 'wf3',
              stage_name: 'In Assessment',
              sla_value: '24',
              sla_unit: 'hours',
              is_active: false
            }
          ]
        },
        {
          id: 'rev2',
          process_name: 'Supplier Assessment',
          customer_id: '1001',
          customer_name: 'PT Suka Suka',
          stage_flow: '06/04/02',
          status: '2 of 3',
          workflows: []
        },
        {
          id: 'rev3',
          process_name: 'Block / Unblock Supplier',
          customer_id: '1002',
          customer_name: 'PT Angin Ribut',
          stage_flow: '01-04-06',
          status: '04: Accounting',
          workflows: []
        }
      ];
    },
    searchProcesses() {
      // Search is handled by computed property
    },
    viewProcess(process) {
      this.selectedProcess = process;
      this.approvalNotes = '';
    },
    closeModal() {
      this.selectedProcess = null;
      this.approvalNotes = '';
    },
    async approveProcess(process) {
      try {
        // In a real implementation, this would call the API
        console.log('Approving process:', process.id);
        process.status = 'Approved';
      } catch (error) {
        console.error('Error approving process:', error);
      }
    },
    async rejectProcess(process) {
      const reason = prompt('Enter reason for rejection:');
      if (reason) {
        try {
          // In a real implementation, this would call the API
          console.log('Rejecting process:', process.id, 'Reason:', reason);
          process.status = 'Rejected';
        } catch (error) {
          console.error('Error rejecting process:', error);
        }
      }
    },
    async submitApproval(action) {
      try {
        // In a real implementation, this would call the API
        console.log('Submitting approval:', action, 'Notes:', this.approvalNotes);
        
        if (this.selectedProcess) {
          switch (action) {
            case 'approved':
              this.selectedProcess.status = 'Approved';
              break;
            case 'rejected':
              this.selectedProcess.status = 'Rejected';
              break;
            case 'returned':
              this.selectedProcess.status = 'Returned';
              break;
          }
        }
        
        this.closeModal();
      } catch (error) {
        console.error('Error submitting approval:', error);
      }
    },
    getStatusClass(status) {
      if (status.includes('Procurement') || status.includes('Approved')) {
        return 'status-active';
      } else if (status.includes('Accounting') || status.includes('Progress')) {
        return 'status-progress';
      } else if (status.includes('Rejected') || status.includes('Blocked')) {
        return 'status-blocked';
      }
      return 'status-pending';
    }
  }
}
</script>

<style scoped>
.review-approvals {
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

.process-tabs {
  display: flex;
  background: white;
  border-radius: 8px 8px 0 0;
  overflow: hidden;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  margin-bottom: 0;
}

.process-tab {
  padding: 15px 25px;
  border: none;
  background: #f8f9fa;
  cursor: pointer;
  font-weight: 500;
  color: #666;
  border-bottom: 3px solid transparent;
  flex: 1;
}

.process-tab.active {
  background: white;
  color: #007bff;
  border-bottom-color: #007bff;
}

.filters {
  background: white;
  padding: 20px;
  border-radius: 0;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  margin-bottom: 0;
}

.search-input {
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
  width: 300px;
}

.table-container {
  background: white;
  border-radius: 0 0 8px 8px;
  overflow: hidden;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.processes-table {
  width: 100%;
  border-collapse: collapse;
}

.processes-table th {
  background-color: #f8f9fa;
  padding: 12px;
  text-align: left;
  font-weight: 600;
  color: #333;
  border-bottom: 1px solid #dee2e6;
}

.processes-table td {
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

.status-pending {
  background-color: #e2e3e5;
  color: #383d41;
}

.status-inactive {
  background-color: #f8f9fa;
  color: #6c757d;
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

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  border-radius: 8px;
  width: 90%;
  max-width: 600px;
  max-height: 90vh;
  overflow-y: auto;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  border-bottom: 1px solid #eee;
}

.modal-header h2 {
  margin: 0;
  color: #333;
}

.close-btn {
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
  color: #666;
}

.modal-body {
  padding: 20px;
}

.process-info {
  margin-bottom: 30px;
}

.info-row {
  display: flex;
  margin-bottom: 10px;
}

.info-row label {
  font-weight: 500;
  width: 120px;
  color: #333;
}

.workflow-section {
  margin-bottom: 30px;
}

.workflow-section h3 {
  color: #333;
  margin-bottom: 15px;
}

.workflow-steps {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.workflow-step {
  background: #f8f9fa;
  padding: 15px;
  border-radius: 6px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.step-name {
  font-weight: 500;
  color: #333;
}

.step-sla {
  font-size: 12px;
  color: #666;
}

.approval-section h3 {
  color: #333;
  margin-bottom: 15px;
}

.approval-notes {
  width: 100%;
  min-height: 100px;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  resize: vertical;
  margin-bottom: 15px;
}

.approval-actions {
  display: flex;
  gap: 10px;
}
</style>


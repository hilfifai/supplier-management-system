document.addEventListener('DOMContentLoaded', function() {
  
    var tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'));
    var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
        return new bootstrap.Tooltip(tooltipTriggerEl);
    });

  
    const sidebarLinks = document.querySelectorAll('.sidebar-menu .nav-link');
    const contentSections = document.querySelectorAll('.content-section');

    sidebarLinks.forEach(link => {
        link.addEventListener('click', function(e) {
            e.preventDefault();
            
          
            sidebarLinks.forEach(l => l.classList.remove('active'));
            
            
            this.classList.add('active');
            
          
            contentSections.forEach(section => {
                section.style.display = 'none';
            });
            
            
            const sectionId = this.getAttribute('data-section') + '-section';
            document.getElementById(sectionId).style.display = 'block';
        });
    });

    
    const tabLinks = document.querySelectorAll('[data-tab]');
    tabLinks.forEach(link => {
        link.addEventListener('click', function(e) {
            e.preventDefault();
            
            
            const parentTabs = this.closest('.nav-tabs');
            parentTabs.querySelectorAll('.nav-link').forEach(tab => {
                tab.classList.remove('active');
            });
            
           
            this.classList.add('active');
            
            
            const tabContent = this.closest('.card-body').querySelectorAll('.tab-content');
            tabContent.forEach(content => {
                content.style.display = 'none';
            });
            
            
            const tabId = this.getAttribute('data-tab') + '-tab';
            document.getElementById(tabId).style.display = 'block';
        });
    });

   
    const newSupplierBtn = document.getElementById('btn-new-supplier');
    const newSupplierModal = new bootstrap.Modal(document.getElementById('newSupplierModal'));
    
    if (newSupplierBtn) {
        newSupplierBtn.addEventListener('click', function() {
            newSupplierModal.show();
        });
    }

    
    const addSupplierBtn = document.getElementById('btn-add-supplier');
    if (addSupplierBtn) {
        addSupplierBtn.addEventListener('click', function() {
            newSupplierModal.show();
        });
    }

    
    const blockSupplierModal = new bootstrap.Modal(document.getElementById('blockSupplierModal'));
    document.querySelectorAll('.btn-block-supplier').forEach(btn => {
        btn.addEventListener('click', function() {
            blockSupplierModal.show();
        });
    });

    
    const confirmBlockBtn = document.getElementById('confirmBlock');
    if (confirmBlockBtn) {
        confirmBlockBtn.addEventListener('click', function() {
            const form = document.getElementById('blockSupplierForm');
            if (!form.checkValidity()) {
                form.classList.add('was-validated');
                return;
            }

            const formData = {
                reason: document.getElementById('blockReason').value,
                
            };

          
            const supplierId = '123'; 

            fetch(`http://localhost:8080/api/v1/suppliers/${supplierId}/block`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData),
            })
            .then(response => response.json())
            .then(data => {
                console.log('Success:', data);
                blockSupplierModal.hide();
                
                fetchSuppliers();
            })
            .catch((error) => {
                console.error('Error:', error);
            });
        });
    }

    
    function fetchDashboardStats() {
        fetch('http://localhost:8080/api/v1/dashboard/stats')
            .then(response => response.json())
            .then(data => {
                document.getElementById('total-suppliers').textContent = data.data.total_suppliers.toLocaleString();
                document.getElementById('new-suppliers').textContent = data.data.new_suppliers;
                document.getElementById('avg-cost').textContent = `Rp ${(data.data.avg_cost_per_supplier / 1000000).toFixed(1)} Mn`;
                document.getElementById('blocked-suppliers').textContent = data.data.blocked_suppliers;
            })
            .catch(error => {
                console.error('Error fetching dashboard stats:', error);
            });
    }

    
    function fetchSuppliers() {
        fetch('http://localhost:8080/api/v1/suppliers')
            .then(response => response.json())
            .then(data => {
                const tableBody = document.querySelector('#supplier-table tbody');
                tableBody.innerHTML = '';

                const table = $('#supplier-table').DataTable({
                    destroy: true,
                    data: data.data, 
                    columns: [
                        { data: 'name', title: 'Name' },
                        { data: 'id', title: 'ID' },
                        { 
                            data: 'status', 
                            title: 'Status',
                            render: function(data, type, row) {
                                return `<span class="badge ${data === 'Active' ? 'bg-success' : 'bg-danger'}">${data}</span>`;
                            }
                        },
                        {
                            title: 'Actions',
                            render: function(data, type, row) {
                                return `
                                    <button class="btn btn-sm btn-primary">View</button>
                                    ${row.status === 'Active' ? 
                                        '<button class="btn btn-sm btn-danger btn-block-supplier">Block</button>' : 
                                        '<button class="btn btn-sm btn-success btn-unblock-supplier">Unblock</button>'}
                                `;
                            },
                            orderable: false 
                        }
                    ],
                    responsive: true, 
                    paging: true, 
                    pageLength: 10 
                });

              
                $('#supplier-table').on('click', '.btn-block-supplier', function() {
                    blockSupplierModal.show();
                });
                
                $('#supplier-table').on('click', '.btn-unblock-supplier', function() {
                    unblockSupplierModal.show();
                });
            })
            .catch(error => {
                console.error('Error fetching suppliers:', error);
            });
    }
    
    function initPage() {
        document.querySelector('.sidebar-menu .nav-link[data-section="dashboard"]').classList.add('active');
        document.getElementById('dashboard-section').style.display = 'block';
        
        fetchDashboardStats();
        fetchSuppliers();
    }

    initPage();
});
document.addEventListener('DOMContentLoaded', function() {
  
    let addressCounter = 1;
    let contactCounter = 1;
    let groupCounter = 1;
    let otherInfoCounter = 1;
    
    
    const supplierForm = document.getElementById('supplierForm');
    const saveSupplierBtn = document.getElementById('saveSupplier');
    const supplierLogoInput = document.getElementById('supplierLogo');
    const supplierLogoPreview = document.getElementById('supplierLogoPreview');
    
   
    initForm();
    
    
    document.getElementById('addAddress').addEventListener('click', addAddressRow);
    document.getElementById('addContact').addEventListener('click', addContactRow);
    document.getElementById('addGroup').addEventListener('click', addGroupRow);
    supplierLogoInput.addEventListener('change', handleLogoUpload);
    saveSupplierBtn.addEventListener('click', saveSupplier);
    
   
    function initForm() {
       
    }
    
    function handleLogoUpload(event) {
        const file = event.target.files[0];
        if (file) {
            const reader = new FileReader();
            reader.onload = function(e) {
                supplierLogoPreview.src = e.target.result;
            };
            reader.readAsDataURL(file);
        }
    }
    
   
    function addAddressRow(data = {}) {
        const tbody = document.querySelector('#addressTable tbody');
        const row = document.createElement('tr');
        const rowId = addressCounter++;
        
        row.innerHTML = `
            <td>${rowId}</td>
            <td>
                <select class="form-select address-type">
                    <option value="Head Office" ${data.type === 'Head Office' ? 'selected' : ''}>Head Office</option>
                    <option value="Branch Office" ${data.type === 'Branch Office' ? 'selected' : ''}>Branch Office</option>
                    <option value="Warehouse" ${data.type === 'Warehouse' ? 'selected' : ''}>Warehouse</option>
                </select>
            </td>
            <td><input type="text" class="form-control address-street" value="${data.street || ''}" required></td>
            <td><input type="text" class="form-control address-city" value="${data.city || ''}" required></td>
            <td><input type="text" class="form-control address-province" value="${data.province || ''}" required></td>
            <td><input type="text" class="form-control address-postal" value="${data.postal_code || ''}"></td>
            <td class="text-center">
                <input type="radio" name="mainAddress" class="form-check-input address-main" ${data.is_main ? 'checked' : ''}>
            </td>
            <td class="text-center">
                <button type="button" class="btn btn-sm btn-danger delete-row">
                    <i class="bi bi-trash"></i>
                </button>
            </td>
        `;
        
        tbody.appendChild(row);
        
        
        row.querySelector('.delete-row').addEventListener('click', () => {
            row.remove();
            renumberTableRows('#addressTable');
        });
        
        row.querySelector('.address-main').addEventListener('change', function() {
            if (this.checked) {
                document.querySelectorAll('.address-main').forEach(radio => {
                    if (radio !== this) radio.checked = false;
                });
            }
        });
    }
    
    
    function addContactRow(data = {}) {
        const tbody = document.querySelector('#contactsTable tbody');
        const row = document.createElement('tr');
        const rowId = contactCounter++;
        
        row.innerHTML = `
            <td>${rowId}</td>
            <td><input type="text" class="form-control contact-name" value="${data.name || ''}" required></td>
            <td><input type="text" class="form-control contact-position" value="${data.job_position || ''}"></td>
            <td><input type="email" class="form-control contact-email" value="${data.email || ''}"></td>
            <td><input type="tel" class="form-control contact-phone" value="${data.phone || ''}"></td>
            <td><input type="tel" class="form-control contact-mobile" value="${data.mobile || ''}"></td>
            <td class="text-center">
                <input type="radio" name="mainContact" class="form-check-input contact-main" ${data.is_main ? 'checked' : ''}>
            </td>
            <td class="text-center">
                <button type="button" class="btn btn-sm btn-danger delete-row">
                    <i class="bi bi-trash"></i>
                </button>
            </td>
        `;
        
        tbody.appendChild(row);
        
        
        row.querySelector('.delete-row').addEventListener('click', () => {
            row.remove();
            renumberTableRows('#contactsTable');
        });
        
        row.querySelector('.contact-main').addEventListener('change', function() {
            if (this.checked) {
                document.querySelectorAll('.contact-main').forEach(radio => {
                    if (radio !== this) radio.checked = false;
                });
            }
        });
    }
    
   
    function addGroupRow(data = {}) {
        const tbody = document.querySelector('#groupsTable tbody');
        const row = document.createElement('tr');
        const rowId = groupCounter++;
        
        row.innerHTML = `
            <td>${rowId}</td>
            <td><input type="text" class="form-control group-name" value="${data.group_name || ''}"></td>
            <td><input type="text" class="form-control group-value" value="${data.value || ''}"></td>
            <td class="text-center">
                <input type="checkbox" class="form-check-input group-active" ${data.is_active ? 'checked' : ''}>
            </td>
            <td class="text-center">
                <button type="button" class="btn btn-sm btn-danger delete-row">
                    <i class="bi bi-trash"></i>
                </button>
            </td>
        `;
        
        tbody.appendChild(row);
        
        
        row.querySelector('.delete-row').addEventListener('click', () => {
            row.remove();
            renumberTableRows('#groupsTable');
        });
    }
    
    
    function renumberTableRows(tableSelector) {
        const rows = document.querySelectorAll(`${tableSelector} tbody tr`);
        rows.forEach((row, index) => {
            row.cells[0].textContent = index + 1;
        });
    }
    
    
    async function saveSupplier() {
        
        if (!supplierForm.checkValidity()) {
            supplierForm.classList.add('was-validated');
            return;
        }
        
       
        const formData = {
            name: document.getElementById('supplierName').value,
            nick_name: document.getElementById('supplierNickname').value,
            category: document.getElementById('supplierCategory').value,
            status: 'Active', 
            addresses: getAddresses(),
            contacts: getContacts(),
            groups: getGroups(),
            other_informations: getOtherInformations()
        };
        console.log(formData);
        
        // Handle logo upload if exists
        // const logoFile = supplierLogoInput.files[0];
        // if (logoFile) {
        //     try {
        //         const logoUrl = await uploadLogo(logoFile);
        //         formData.logo_url = logoUrl;
        //     } catch (error) {
        //         showToast('Error uploading logo: ' + error.message, 'danger');
        //         return;
        //     }
        // }
       
        try {
            const response = await fetch('http://localhost:8080/api/v1/suppliers', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'X-CSRF-TOKEN': document.querySelector('meta[name="csrf-token"]').content
                },
                body: JSON.stringify(formData)
            });
            
            if (!response.ok) {
                throw new Error(await response.text());
            }
            
            const data = await response.json();
            showToast('Supplier created successfully!', 'success');
            
            bootstrap.Modal.getInstance(document.getElementById('newSupplierModal')).hide();
           
        } catch (error) {
            console.error('Error:', error);
            showToast('Error creating supplier: ' + error.message, 'danger');
        }
    }
    
   
    async function uploadLogo(file) {
        const formData = new FormData();
        formData.append('logo', file);
        
        const response = await fetch('http://localhost:8080/api/v1/upload-logo', {
            method: 'POST',
            headers: {
                'X-CSRF-TOKEN': document.querySelector('meta[name="csrf-token"]').content
            },
            body: formData
        });
        
        if (!response.ok) {
            throw new Error('Logo upload failed');
        }
        
        const data = await response.json();
        return data.url;
    }
    
   
    function getAddresses() {
        const addresses = [];
        document.querySelectorAll('#addressTable tbody tr').forEach(row => {
            addresses.push({
                type: row.querySelector('.address-type').value,
                street: row.querySelector('.address-street').value,
                city: row.querySelector('.address-city').value,
                province: row.querySelector('.address-province').value,
                postal_code: row.querySelector('.address-postal').value,
                is_main: row.querySelector('.address-main').checked
            });
        });
        return addresses;
    }
    
    function getContacts() {
        const contacts = [];
        document.querySelectorAll('#contactsTable tbody tr').forEach(row => {
            console.log(row);
            contacts.push({
                name: row.querySelector('.contact-name').value,
                job_position: row.querySelector('.contact-position').value,
                email: row.querySelector('.contact-email').value,
                phone: row.querySelector('.contact-phone').value,
                mobile: row.querySelector('.contact-mobile').value,
                is_main: row.querySelector('.contact-main').checked
            });
        });
        return contacts;
    }
    
    function getGroups() {
        const groups = [];
        document.querySelectorAll('#groupsTable tbody tr').forEach(row => {
            groups.push({
                group_name: row.querySelector('.group-name').value,
                value: row.querySelector('.group-value').value,
                is_active: row.querySelector('.group-active').checked
            });
        });
        return groups;
    }
    
    function getOtherInformations() {
       
        return [];
    }
    
   
    function showToast(message, type = 'success') {
        const toastContainer = document.getElementById('toastContainer') || createToastContainer();
        const toast = document.createElement('div');
        toast.className = `toast show align-items-center text-white bg-${type} border-0`;
        toast.setAttribute('role', 'alert');
        toast.innerHTML = `
            <div class="d-flex">
                <div class="toast-body">${message}</div>
                <button type="button" class="btn-close btn-close-white me-2 m-auto" data-bs-dismiss="toast"></button>
            </div>
        `;
        toastContainer.appendChild(toast);
        
        setTimeout(() => toast.remove(), 5000);
    }
    
    function createToastContainer() {
        const container = document.createElement('div');
        container.id = 'toastContainer';
        container.className = 'toast-container position-fixed bottom-0 end-0 p-3';
        document.body.appendChild(container);
        return container;
    }
});
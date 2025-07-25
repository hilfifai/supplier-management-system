<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>ALISA - Supplier Management</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.0/font/bootstrap-icons.css">
    <link href="{{ asset('css/supplier.css') }}" rel="stylesheet">

    <meta name="csrf-token" content="{{ csrf_token() }}">
</head>

<body>
    <div class="container-fluid">
        <div class="row">
            <div class="col-md-2 sidebar">
                <div class="sidebar-header">
                    <h1>ALISA</h1>
                </div>
                <div class="sidebar-menu">
                    <h5>Dashboard</h5>
                    <ul class="nav flex-column">
                        <li class="nav-item">
                            <a class="nav-link active" href="#" data-section="dashboard">
                                <i class="bi bi-speedometer2"></i> Dashboard
                            </a>
                        </li>
                    </ul>
                     <h5>Supplier Management </h5>
                     <ul class="nav flex-column">
                        <li class="nav-item">
                            <a class="nav-link" href="#" data-section="supplier-list">
                                <i class="bi bi-list-ul"></i> Supplier List
                            </a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="#" data-section="review-approvals">
                                <i class="bi bi-clipboard-check"></i> Review & Approvals
                            </a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="#" data-section="configurations">
                                <i class="bi bi-gear"></i> Configurations
                            </a>
                        </li>
                    </ul>

                    <h5>Funnel Management</h5>
                    <ul class="nav flex-column">
                        <li class="nav-item">
                            <a class="nav-link" href="#" data-section="funnel-management">
                                <i class="bi bi-funnel"></i> Funnel Overview
                            </a>
                        </li>
                    </ul>
                </div>
                <div class="sidebar-footer">
                    <div class="help-support">
                        <i class="bi bi-question-circle"></i> Help & Support: Hilfi Muhamad Aryawan
                    </div>
                </div>
            </div>

            <!-- Main Content -->
            <div class="col-md-10 main-content">
                <!-- Dashboard Section -->
                <section id="dashboard-section" class="content-section" style="display: none;">
                    <div class="row">
                        <div class="col-md-12">
                            <h2>Dashboard</h2>
                            <p>Welcome to the Supplier Management Dashboard. Here you can manage suppliers, review approvals, and configure settings.</p>
                        </div>
                    </div>
                </section>
                <section id="supplier-list-section" class="content-section">
                    <div class="row mb-4">
                        <div class="col-md-12">
                            <h2>Supplier List</h2>
                            <div class="dashboard-stats">
                                <div class="row">
                                    <div class="col-md-3">
                                        <div class="stat-card">
                                            <h5>Total Supplier</h5>
                                            <h3 id="total-suppliers">1,869</h3>
                                            <p class="text-success">+8% vs last year</p>
                                        </div>
                                    </div>
                                    <div class="col-md-3">
                                        <div class="stat-card">
                                            <h5>New Supplier</h5>
                                            <h3 id="new-suppliers">27</h3>
                                            <p class="text-success">+7% vs Last Year</p>
                                        </div>
                                    </div>
                                    <div class="col-md-3">
                                        <div class="stat-card">
                                            <h5>Avg Cost per Supplier</h5>
                                            <h3 id="avg-cost">Rp 320,3 Mn</h3>
                                            <p class="text-success">+5% vs Last Year</p>
                                        </div>
                                    </div>
                                    <div class="col-md-3">
                                        <div class="stat-card">
                                            <h5>Blocked Supplier</h5>
                                            <h3 id="blocked-suppliers">31</h3>
                                            <p class="text-success">+5% vs Last Year</p>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="row">
                        <div class="col-md-12">
                            <div class="card">
                                <div class="card-header">
                                    <h5>Supplier List</h5>
                                    <button class="btn btn-primary mb-2" id="btn-new-supplier">
                                        <i class="bi bi-plus-circle"></i> New Supplier
                                    </button>
                                    <button class="btn btn-secondary mb-2" id="btn-export">
                                        <i class="bi bi-download"></i> Export Data
                                    </button>
                                </div>
                                <div class="card-body">
                                    <table class="table table-hover" id="supplier-table" style="width: 100%">
                                        <thead>
                                            <tr>
                                                <th>Name</th>
                                                <th>ID</th>
                                                <th>Status</th>
                                                <th>Actions</th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            
                                        </tbody>
                                    </table>
                                    <div class="d-flex justify-content-between">
                                       
                                       
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="col-md-4">
                           
                            
                            </div>
                        </div>
                    </div>
                </section>

                <!-- Supplier List Section -->
                <section id="supplier-list-section2" class="content-section" style="display: none;">
                    <div class="row">
                        <div class="col-md-12">
                            <h2>Supplier List</h2>
                            <div class="card">
                                <div class="card-header">
                                    <div class="row">
                                        <div class="col-md-6">
                                            <div class="search-box">
                                                <input type="text" id="supplier-list-search" placeholder="Search supplier...">
                                                <button class="btn btn-sm btn-primary">Search</button>
                                            </div>
                                        </div>
                                        <div class="col-md-6 text-end">
                                            <button class="btn btn-primary" id="btn-add-supplier">
                                                <i class="bi bi-plus"></i> Add Supplier
                                            </button>
                                        </div>
                                    </div>
                                </div>
                                <div class="card-body">
                                    <table class="table table-hover" id="detailed-supplier-table">
                                        <thead>
                                            <tr>
                                                <th>Name</th>
                                                <th>ID</th>
                                                <th>Nickname</th>
                                                <th>Address</th>
                                                <th>Contacts</th>
                                                <th>Groups</th>
                                                <th>Status</th>
                                                <th>Actions</th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            
                                        </tbody>
                                    </table>
                                    <div class="d-flex justify-content-between">
                                        
                                       
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </section>

                <!-- Review & Approvals Section -->
                <section id="review-approvals-section" class="content-section" style="display: none;">
                    <div class="row">
                        <div class="col-md-12">
                            <h2>Review & Approvals</h2>
                            <div class="card">
                                <div class="card-header">
                                    <ul class="nav nav-tabs card-header-tabs">
                                        <li class="nav-item">
                                            <a class="nav-link active" href="#" data-tab="supplier-creation">Supplier Creation</a>
                                        </li>
                                        <li class="nav-item">
                                            <a class="nav-link" href="#" data-tab="supplier-assessment">Supplier Assessment</a>
                                        </li>
                                        <li class="nav-item">
                                            <a class="nav-link" href="#" data-tab="block-unblock">Block / Unblock Supplier</a>
                                        </li>
                                    </ul>
                                </div>
                                <div class="card-body">
                                    <div id="supplier-creation-tab" class="tab-content">
                                        <div class="search-box mb-3">
                                            <input type="text" id="search-supplier-creation" placeholder="Search process...">
                                            <button class="btn btn-sm btn-primary">Search</button>
                                        </div>
                                        <table class="table table-hover">
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
                                                <tr>
                                                    <td>Supplier Creation</td>
                                                    <td>1001</td>
                                                    <td>PT Suko Suko</td>
                                                    <td>01-02-03-05</td>
                                                    <td>02: Procurement</td>
                                                    <td>
                                                        <button class="btn btn-sm btn-primary">Review</button>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </div>
                                    <div id="supplier-assessment-tab" class="tab-content" style="display: none;">
                                        <!-- Supplier Assessment content -->
                                    </div>
                                    <div id="block-unblock-tab" class="tab-content" style="display: none;">
                                        <!-- Block/Unblock content -->
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </section>

                <!-- Configurations Section -->
                <section id="configurations-section" class="content-section" style="display: none;">
                    <div class="row">
                        <div class="col-md-12">
                            <h2>Configurations</h2>
                            <div class="card">
                                <div class="card-header">
                                    <ul class="nav nav-tabs card-header-tabs">
                                        <li class="nav-item">
                                            <a class="nav-link active" href="#" data-tab="workflows">Workflows</a>
                                        </li>
                                        <li class="nav-item">
                                            <a class="nav-link" href="#" data-tab="stages">Stages</a>
                                        </li>
                                        <li class="nav-item">
                                            <a class="nav-link" href="#" data-tab="others">Others</a>
                                        </li>
                                    </ul>
                                </div>
                                <div class="card-body">
                                    <div id="workflows-tab" class="tab-content">
                                        <h5>Workflow Configuration</h5>
                                        <p>Configure your supplier management workflows here.</p>
                                    </div>
                                    <div id="stages-tab" class="tab-content" style="display: none;">
                                        <table class="table table-hover">
                                            <thead>
                                                <tr>
                                                    <th>ID</th>
                                                    <th>Name</th>
                                                    <th>SLA</th>
                                                    <th>UoM</th>
                                                    <th>Active</th>
                                                    <th>Actions</th>
                                                </tr>
                                            </thead>
                                            <tbody>
                                                <tr>
                                                    <td>01</td>
                                                    <td>Draft</td>
                                                    <td>24</td>
                                                    <td>hours</td>
                                                    <td><span class="badge bg-success">Yes</span></td>
                                                    <td>
                                                        <button class="btn btn-sm btn-primary">Edit</button>
                                                    </td>
                                                </tr>
                                                <tr>
                                                    <td>02</td>
                                                    <td>In Review</td>
                                                    <td>24</td>
                                                    <td>hours</td>
                                                    <td><span class="badge bg-success">Yes</span></td>
                                                    <td>
                                                        <button class="btn btn-sm btn-primary">Edit</button>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </div>
                                    <div id="others-tab" class="tab-content" style="display: none;">
                                        <!-- Others content -->
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </section>

                <!-- Funnel Management Section -->
                <section id="funnel-management-section" class="content-section" style="display: none;">
                    <div class="row">
                        <div class="col-md-12">
                            <h2>Funnel Management</h2>
                            <div class="card">
                                <div class="card-header">
                                    <h5>Supplier Funnel Overview</h5>
                                </div>
                                <div class="card-body">
                                    <div class="funnel-chart">
                                        <!-- Placeholder for funnel chart -->
                                        <img src="https://via.placeholder.com/800x400?text=Funnel+Chart" alt="Funnel Chart" class="img-fluid">
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </section>

                <!-- Modal for New Supplier -->
                <div class="modal fade" id="newSupplierModal" tabindex="-1" aria-labelledby="newSupplierModalLabel" aria-hidden="true">
                    <div class="modal-dialog modal-lg">
                        <div class="modal-content">
                            <div class="modal-header">
                                <h5 class="modal-title" id="newSupplierModalLabel">New Supplier</h5>
                                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                            </div>
                            <div class="modal-body">
                                <form id="supplierForm">
                                    <div class="row mb-3">
                                        <div class="col-md-2">
                                            <label class="form-label">Logo</label>
                                            <div class="logo-upload">
                                                <img src="https://via.placeholder.com/100" alt="Supplier Logo" id="supplierLogoPreview" class="img-thumbnail">
                                                <input type="file" id="supplierLogo" class="d-none">
                                                <button type="button" class="btn btn-sm btn-secondary mt-2" onclick="document.getElementById('supplierLogo').click()">Upload</button>
                                            </div>
                                        </div>
                                        <div class="col-md-5">
                                            <div class="mb-3">
                                                <label for="supplierName" class="form-label">Supplier Name</label>
                                                <input type="text" class="form-control" id="supplierName" required>
                                            </div>
                                            <div class="mb-3">
                                                <label for="supplierNickname" class="form-label">Nick Name</label>
                                                <input type="text" class="form-control" id="supplierNickname">
                                            </div>
                                        </div>
                                        <div class="col-md-5">
                                            <div class="mb-3">
                                                <label for="supplierCategory" class="form-label">Supplier Category</label>
                                                <select class="form-select" id="supplierCategory">
                                                    <option value="manufacture">Manufacture</option>
                                                    <option value="service">Service</option>
                                                    <option value="retail">Retail</option>
                                                </select>
                                            </div>
                                        </div>
                                    </div>

                                    <ul class="nav nav-tabs" id="supplierTabs" role="tablist">
                                        <li class="nav-item" role="presentation">
                                            <button class="nav-link active" id="address-tab" data-bs-toggle="tab" data-bs-target="#address" type="button" role="tab">Address</button>
                                        </li>
                                        <li class="nav-item" role="presentation">
                                            <button class="nav-link" id="contacts-tab" data-bs-toggle="tab" data-bs-target="#contacts" type="button" role="tab">Contacts</button>
                                        </li>
                                        <li class="nav-item" role="presentation">
                                            <button class="nav-link" id="groups-tab" data-bs-toggle="tab" data-bs-target="#groups" type="button" role="tab">Groups</button>
                                        </li>
                                        <li class="nav-item" role="presentation">
                                            <button class="nav-link" id="materials-tab" data-bs-toggle="tab" data-bs-target="#materials" type="button" role="tab">Material List</button>
                                        </li>
                                        <li class="nav-item" role="presentation">
                                            <button class="nav-link" id="others-tab" data-bs-toggle="tab" data-bs-target="#others" type="button" role="tab">Others</button>
                                        </li>
                                    </ul>

                                    <div class="tab-content p-3 border border-top-0" id="supplierTabContent">
                                        <div class="tab-pane fade show active" id="address" role="tabpanel">
                                            <div class="mb-3">
                                                <button class="btn btn-sm btn-primary" id="addAddress">Add Address</button>
                                            </div>
                                            <table class="table table-bordered" id="addressTable">
                                                <thead>
                                                    <tr>
                                                        <th>#</th>
                                                        <th>Name</th>
                                                        <th>Address</th>
                                                        <th>City</th>
                                                        <th>Province</th>
                                                        <th>Postal</th>
                                                        <th>Main</th>
                                                        <th>Actions</th>
                                                    </tr>
                                                </thead>
                                                <tbody>

                                                </tbody>
                                            </table>
                                        </div>
                                        <div class="tab-pane fade" id="contacts" role="tabpanel">
                                            <div class="mb-3">
                                                <button class="btn btn-sm btn-primary" id="addContact">Add Contact</button>
                                            </div>
                                            <table class="table table-bordered" id="contactsTable">
                                                <thead>
                                                    <tr>
                                                        <th>#</th>
                                                        <th>Name</th>
                                                        <th>Job Position</th>
                                                        <th>Email</th>
                                                        <th>Phone</th>
                                                        <th>Mobile</th>
                                                        <th>Main</th>
                                                        <th>Actions</th>
                                                    </tr>
                                                </thead>
                                                <tbody>

                                                </tbody>
                                            </table>
                                        </div>
                                        <div class="tab-pane fade" id="groups" role="tabpanel">
                                            <div class="mb-3">
                                                <button class="btn btn-sm btn-primary" id="addGroup">Add Group</button>
                                            </div>
                                            <table class="table table-bordered" id="groupsTable">
                                                <thead>
                                                    <tr>
                                                        <th>#</th>
                                                        <th>Group Name</th>
                                                        <th>Value</th>
                                                        <th>Active</th>
                                                        <th>Actions</th>
                                                    </tr>
                                                </thead>
                                                <tbody>

                                                </tbody>
                                            </table>
                                        </div>
                                        <div class="tab-pane fade" id="materials" role="tabpanel">
                                            <!-- Material List content -->
                                        </div>
                                        <div class="tab-pane fade" id="others" role="tabpanel">
                                            <!-- Others content -->
                                        </div>
                                    </div>
                                </form>
                            </div>
                            <div class="modal-footer">
                                <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cancel</button>
                                <button type="button" class="btn btn-primary" id="saveSupplier">Save</button>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- Modal for Block Supplier -->
                <div class="modal fade" id="blockSupplierModal" tabindex="-1" aria-labelledby="blockSupplierModalLabel" aria-hidden="true">
                    <div class="modal-dialog">
                        <div class="modal-content">
                            <div class="modal-header">
                                <h5 class="modal-title" id="blockSupplierModalLabel">Block Supplier</h5>
                                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                            </div>
                            <div class="modal-body">
                                <form id="blockSupplierForm">
                                    <div class="mb-3">
                                        <label for="blockReason" class="form-label">Reason</label>
                                        <textarea class="form-control" id="blockReason" rows="3" required></textarea>
                                    </div>
                                    <div class="mb-3">
                                        <label for="blockDocuments" class="form-label">Documents</label>
                                        <input type="file" class="form-control" id="blockDocuments" multiple>
                                    </div>
                                </form>
                            </div>
                            <div class="modal-footer">
                                <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cancel</button>
                                <button type="button" class="btn btn-danger" id="confirmBlock">Block Supplier</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
    
    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>

    <link rel="stylesheet" type="text/css" href="https://cdn.datatables.net/1.11.5/css/jquery.dataTables.css">

    <script type="text/javascript" charset="utf8" src="https://cdn.datatables.net/1.11.5/js/jquery.dataTables.js"></script>

    <link rel="stylesheet" type="text/css" href="https://cdn.datatables.net/1.11.5/css/dataTables.bootstrap5.min.css">
    <script src="https://cdn.datatables.net/1.11.5/js/dataTables.bootstrap5.min.js"></script>
    <script src="{{ asset('js/supplier.js') }}"></script>
</body>

</html>

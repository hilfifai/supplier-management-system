import { createApp } from 'vue';
import axios from 'axios';
import Dashboard from './components/Dashboard.vue';
import SupplierList from './components/SupplierList.vue';
import SupplierDetail from './components/SupplierDetail.vue';
import NewSupplier from './components/NewSupplier.vue';
import ReviewApprovals from './components/ReviewApprovals.vue';

axios.defaults.baseURL = 'http://localhost:8080/api/v1';

const app = createApp({
    data() {
        return {
            currentView: 'dashboard',
            selectedSupplier: null
        }
    },
    methods: {
        navigateTo(view, data = null) {
            this.currentView = view;
            if (data) {
                this.selectedSupplier = data;
            }
        }
    }
});

app.component('Dashboard', Dashboard);
app.component('SupplierList', SupplierList);
app.component('SupplierDetail', SupplierDetail);
app.component('NewSupplier', NewSupplier);
app.component('ReviewApprovals', ReviewApprovals);

app.mount('#app');


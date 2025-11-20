import { createRouter, createWebHistory } from 'vue-router'

import HomeView from '../views/HomeView.vue';
import ClientesView from '../views/ClientesView.vue';
import ProdutosView from '../views/ProdutosView.vue';
import OfertasView from '../views/OfertasView.vue';
import FornecedoresView from '../views/FornecedoresView.vue';

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
          path: '/',
          name: 'Home',
          component: HomeView,
        },
        {
          path: '/clientes',
          name: 'Clientes',
          component: ClientesView,
        },
        {
          path: '/produtos',
          name: 'Produtos',
          component: ProdutosView,
        },
        {
          path: '/ofertas',
          name: 'Ofertas',
          component: OfertasView,
        },
        {
          path: '/fornecedores',
          name: 'Fornecedores',
          component: FornecedoresView,
        },
    ],
})

export default router

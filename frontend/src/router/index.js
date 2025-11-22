import { createRouter, createWebHistory } from "vue-router";

import HomeView from "../views/HomeView.vue";
import FinanceiroView from "../views/FinanceiroView.vue";
import ProdutosView from "../views/ProdutosView.vue";
import VendasView from "../views/VendasView.vue";
import FornecedoresView from "../views/FornecedoresView.vue";
import FuncionariosView from "../views/FuncionariosView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "Home",
      component: HomeView,
    },
    {
      path: "/financeiro",
      name: "Financeiro",
      component: FinanceiroView,
    },
    {
      path: "/produtos",
      name: "Produtos",
      component: ProdutosView,
    },
    {
      path: "/vendas",
      name: "Vendas",
      component: VendasView,
    },
    {
      path: "/fornecedores",
      name: "Fornecedores",
      component: FornecedoresView,
    },
    {
      path: "/funcionarios",
      name: "Funcionarios",
      component: FuncionariosView,
    },
  ],
});

export default router;

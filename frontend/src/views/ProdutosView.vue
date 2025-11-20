<script setup>
import { ref, onMounted } from 'vue';
import api from '../services/api';

const produtos = ref([]);
const loading = ref(true);
const error = ref(null);

async function fetchProdutos() {
  try {
    loading.value = true;
    const response = await api.getProdutos();
    produtos.value = response.data;
  } catch (err) {
    error.value = 'Falha ao buscar produtos.';
  } finally {
    loading.value = false;
  }
}

onMounted(fetchProdutos);
</script>

<template>
  <div class="db-table">
    <h2>Lista de Produtos</h2>
    
    <div v-if="loading">Carregando...</div>
    <div v-else-if="error" style="color: red">{{ error }}</div>
    
    <table v-else-if="produtos.length > 0">
      <thead>
        <tr>
          <th>ID</th>
          <th>Nome</th>
          <th>Categoria</th>
          <th>Marca</th>
          <th>Preço Venda</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="p in produtos" :key="p.id">
          <td>{{ p.id }}</td>
          <td>{{ p.nome }}</td>
          <td>{{ p.categoria }}</td>
          <td>{{ p.marca }}</td>
          <td>{{ p.preco_venda ? `R$ ${p.preco_venda.toFixed(2)}` : '-' }}</td>
        </tr>
      </tbody>
    </table>
    <div v-else>Nenhum produto encontrado.</div>
  </div>
</template>

<style scoped>
</style>

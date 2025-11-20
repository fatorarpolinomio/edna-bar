<script setup>
import { ref, onMounted } from 'vue';
import api from '../services/api'; // Importa nosso serviço

const fornecedores = ref([]);
const loading = ref(true);
const error = ref(null);

async function fetchFornecedores() {
  try {
    loading.value = true;
    const response = await api.getFornecedores();
    fornecedores.value = response.data;
  } catch (err) {
    error.value = 'Falha ao buscar fornecedores.';
  } finally {
    loading.value = false;
  }
}

onMounted(fetchFornecedores);
</script>

<template>
  <div class="db-table">
    <h2>Lista de Fornecedores</h2>
    
    <div v-if="loading">Carregando...</div>
    <div v-else-if="error" style="color: red">{{ error }}</div>
    
    <table v-else-if="fornecedores.length > 0">
      <thead>
        <tr>
          <th>ID</th>
          <th>Nome</th>
          <th>CNPJ</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="f in fornecedores" :key="f.id">
          <td>{{ f.id }}</td>
          <td>{{ f.nome }}</td>
          <td>{{ f.cnpj }}</td>
        </tr>
      </tbody>
    </table>
    <div v-else>Nenhum fornecedor encontrado.</div>
  </div>
</template>

<style scoped>
</style>

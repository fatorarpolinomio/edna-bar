<script setup>
import { ref, onMounted } from 'vue';
import api from '../services/api';

const clientes = ref([]);
const loading = ref(true);
const error = ref(null);

async function fetchClientes() {
  try {
    loading.value = true;
    const response = await api.getClientes();
    clientes.value = response.data;
  } catch (err) {
    error.value = 'Falha ao buscar clientes.';
  } finally {
    loading.value = false;
  }
}

onMounted(fetchClientes);
</script>

<template>
  <div class="db-table">
    <h2>Lista de Clientes</h2>
    
    <div v-if="loading">Carregando...</div>
    <div v-else-if="error" style="color: red">{{ error }}</div>
    
    <table v-else-if="clientes.length > 0">
      <thead>
        <tr>
          <th>ID</th>
          <th>Nome</th>
          <th>CPF</th>
          <th>Data Nasc.</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="c in clientes" :key="c.id">
          <td>{{ c.id }}</td>
          <td>{{ c.nome }}</td>
          <td>{{ c.cpf || '-' }}</td>
          <td>{{ c.data_nascimento || '-' }}</td>
        </tr>
      </tbody>
    </table>
    <div v-else>Nenhum cliente encontrado.</div>
  </div>
</template>

<style scoped>
</style>

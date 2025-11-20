<script setup>
import { ref, onMounted } from 'vue';
import api from '../services/api';

const ofertas = ref([]);
const loading = ref(true);
const error = ref(null);

async function fetchOfertas() {
  try {
    loading.value = true;
    const response = await api.getOfertas();
    ofertas.value = response.data;
  } catch (err) {
    error.value = 'Falha ao buscar ofertas.';
  } finally {
    loading.value = false;
  }
}

onMounted(fetchOfertas);
</script>

<template>
  <div class="db-table">
    <h2>Lista de Ofertas</h2>
    
    <div v-if="loading">Carregando...</div>
    <div v-else-if="error" style="color: red">{{ error }}</div>
    
    <table v-else-if="ofertas.length > 0">
      <thead>
        <tr>
          <th>ID</th>
          <th>Nome</th>
          <th>Valor Fixo</th>
          <th>Desconto (%)</th>
          <th>Início</th>
          <th>Fim</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="o in ofertas" :key="o.id_oferta">
          <td>{{ o.id_oferta }}</td>
          <td>{{ o.nome }}</td>
          <td>{{ o.valor_fixo ? `R$ ${o.valor_fixo.toFixed(2)}` : '-' }}</td>
          <td>{{ o.percentual_desconto ? `${o.percentual_desconto}%` : '-' }}</td>
          <td>{{ o.data_inicio || '-' }}</td>
          <td>{{ o.data_fim || '-' }}</td>
        </tr>
      </tbody>
    </table>
    <div v-else>Nenhuma oferta encontrada.</div>
  </div>
</template>

<style scoped>
</style>

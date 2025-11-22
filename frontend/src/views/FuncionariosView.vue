<script setup>
import { ref, onMounted, reactive } from 'vue';
import api from '@/services/api';

// --- ESTADO ---
const funcionarios = ref([]);
const itensEstrutura = ref([]);
const carregando = ref(false);

// Formulário Funcionário (Baseado no Model do Swagger/Go)
const formFuncionario = reactive({
  nome: '',
  CPF: '',
  tipo: 'garcom', // Default
  expediente: 'noite', // Default
  salario: '',
  data_contratacao: new Date().toISOString().split('T')[0] // Hoje YYYY-MM-DD
});

// Formulário Item Estrutural
const formItem = reactive({
  nome: '',
  categoria: '',
  marca: ''
});

// --- AÇÕES ---

const carregarDados = async () => {
  carregando.value = true;
  try {
    const [resFunc, resEst] = await Promise.all([
      api.getFuncionarios(),
      api.getProdutosEstruturais()
    ]);

    funcionarios.value = resFunc.data || [];

    // Carregar quantidades dos itens estruturais (opcional, se necessário)
    const itensTemp = resEst.data || [];
    itensEstrutura.value = await Promise.all(itensTemp.map(async (p) => {
      try {
        const resQtd = await api.getProdutoQtd(p.id);
        return { ...p, quantidade: resQtd.data.quantidade_disponível };
      } catch {
        return { ...p, quantidade: 0 };
      }
    }));

  } catch (error) {
    console.error("Erro ao carregar:", error);
  } finally {
    carregando.value = false;
  }
};

const salvarFuncionario = async () => {
  if(!formFuncionario.nome || !formFuncionario.CPF) return alert("Dados incompletos");

  try {
    // Salario deve ser float no Go
    const payload = {
      ...formFuncionario,
      salario: parseFloat(formFuncionario.salario)
    };

    await api.createFuncionario(payload);

    // Resetar form
    Object.assign(formFuncionario, {
      nome: '', CPF: '', salario: '',
      tipo: 'garcom', expediente: 'noite',
      data_contratacao: new Date().toISOString().split('T')[0]
    });

    await carregarDados();
  } catch (e) {
    alert("Erro: " + (e.response?.data?.detail || e.message));
  }
};

const salvarItem = async () => {
  try {
    await api.createProduto(formItem);
    Object.assign(formItem, { nome: '', categoria: '', marca: '' });
    await carregarDados();
  } catch (e) {
    alert("Erro ao criar item.");
  }
};

const deletar = async (contexto, id) => {
  if(!confirm("Remover este registro?")) return;
  try {
    if(contexto === 'func') await api.deleteFuncionario(id);
    if(contexto === 'item') await api.deleteByEndpoint(`/produtos/${id}`);
    await carregarDados();
  } catch (e) {
    alert("Erro ao deletar.");
  }
};

onMounted(() => {
  carregarDados();
});
</script>

<template>
  <div class="nav-space"></div>
  <div class="page-container">

    <section class="card-panel">
      <div class="panel-header">
        <h2>Funcionários</h2>
      </div>

      <div class="form-row">
        <input v-model="formFuncionario.nome" placeholder="Nome Completo" />
        <input v-model="formFuncionario.CPF" placeholder="CPF (somente núm.)" maxlength="11" />

        <select v-model="formFuncionario.tipo">
          <option value="garcom">Garçom</option>
          <option value="seguranca">Segurança</option>
          <option value="caixa">Caixa</option>
          <option value="faxineiro">Faxineiro</option>
          <option value="balconista">Balconista</option>
        </select>

        <select v-model="formFuncionario.expediente">
          <option value="noite">Noite</option>
          <option value="manha">Manhã</option>
          <option value="tarde">Tarde</option>
        </select>

        <input type="number" v-model="formFuncionario.salario" placeholder="Salário R$" class="input-short" />
        <input type="date" v-model="formFuncionario.data_contratacao" />

        <button @click="salvarFuncionario" class="btn-action">+</button>
      </div>

      <div class="table-wrapper">
        <table>
          <thead>
            <tr>
              <th>Nome</th>
              <th>CPF</th>
              <th>Função</th>
              <th>Expediente</th>
              <th>Salário</th>
              <th>Ações</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="func in funcionarios" :key="func.id">
              <td>{{ func.nome }}</td>
              <td>{{ func.CPF }}</td>
              <td><span class="badge">{{ func.tipo }}</span></td>
              <td>{{ func.expediente }}</td>
              <td class="money">R$ {{ func.salario.toFixed(2) }}</td>
              <td>
                <button @click="deletar('func', func.id)" class="btn-delete">Excluir</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </section>

    <section class="card-panel mt-large">
      <div class="panel-header">
        <h2>Itens & Estrutura</h2>
        <p class="subtitle">Gerenciamento de insumos, mobília e descartáveis</p>
      </div>

      <div class="form-row">
        <input v-model="formItem.nome" placeholder="Nome do Item (ex: Copo Descartável)" />
        <input v-model="formItem.categoria" placeholder="Categoria (ex: Descartáveis)" />
        <input v-model="formItem.marca" placeholder="Marca" />
        <button @click="salvarItem" class="btn-action">+</button>
      </div>

      <div class="grid-items">
        <div v-for="item in itensEstrutura" :key="item.id" class="item-card">
          <div class="item-info">
            <h3>{{ item.nome }}</h3>
            <p>{{ item.categoria }} | {{ item.marca }}</p>
          </div>
          <div class="item-qtd">
            <span class="qtd-label">Qtd:</span>
            <span class="qtd-value">{{ item.quantidade }}</span>
          </div>
          <button @click="deletar('item', item.id)" class="btn-close">×</button>
        </div>
      </div>
    </section>

  </div>
</template>

<style scoped>
/* Estilo baseado no Tema E.D.N.A e no desenho */

.page-container {
  background-color: var(--edna-black);
  min-height: 100vh;
  padding: 2rem;
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.card-panel {
  background-color: var(--edna-dark-gray);
  border: 1px solid var(--edna-gray);
  border-radius: 15px;
  padding: 1.5rem;
  box-shadow: 0 4px 10px rgba(0,0,0,0.3);
}

.panel-header h2 {
  color: var(--edna-yellow);
  font-size: 1.8rem;
  margin-bottom: 0.5rem;
  border-bottom: 2px solid var(--edna-wine);
  display: inline-block;
  padding-bottom: 5px;
}

.subtitle {
  color: var(--edna-light-gray);
  font-size: 0.9rem;
  margin-bottom: 1rem;
}

/* --- Formulários --- */
.form-row {
  display: flex;
  gap: 10px;
  flex-wrap: wrap;
  background-color: rgba(255, 255, 255, 0.05);
  padding: 15px;
  border-radius: 8px;
  margin-bottom: 20px;
  align-items: center;
}

input, select {
  background-color: var(--edna-black);
  color: var(--edna-white);
  border: 1px solid var(--edna-gray);
  padding: 10px;
  border-radius: 5px;
  flex: 1;
  min-width: 120px;
}

input:focus, select:focus {
  border-color: var(--edna-green);
  outline: none;
}

.input-short {
  max-width: 100px;
}

.btn-action {
  background-color: var(--edna-green);
  color: var(--edna-black);
  font-weight: bold;
  font-size: 1.5rem;
  width: 45px;
  height: 45px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
}

/* --- Tabela (Funcionários) --- */
.table-wrapper {
  overflow-x: auto;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th {
  background-color: var(--edna-gray);
  color: var(--edna-yellow);
  text-align: left;
  padding: 12px;
}

td {
  border-bottom: 1px solid var(--edna-gray);
  padding: 12px;
  color: var(--edna-white);
}

.badge {
  background-color: var(--edna-wine);
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 0.8rem;
  text-transform: uppercase;
}

.money {
  color: var(--edna-green);
  font-family: monospace;
}

.btn-delete {
  background-color: var(--edna-red);
  color: white;
  padding: 5px 10px;
  border-radius: 4px;
  font-size: 0.8rem;
}

/* --- Cards (Itens) --- */
.grid-items {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 15px;
}

.item-card {
  background-color: var(--edna-gray);
  border-radius: 8px;
  padding: 15px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: relative;
  border: 1px solid transparent;
}

.item-card:hover {
  border-color: var(--edna-blue);
}

.item-info h3 {
  margin: 0;
  font-size: 1.1rem;
  color: var(--edna-white);
}

.item-info p {
  margin: 5px 0 0 0;
  font-size: 0.85rem;
  color: var(--edna-light-gray);
}

.item-qtd {
  text-align: center;
  background-color: var(--edna-black);
  padding: 8px;
  border-radius: 6px;
  min-width: 60px;
}

.qtd-label {
  display: block;
  font-size: 0.7rem;
  color: var(--edna-light-gray);
}

.qtd-value {
  font-size: 1.2rem;
  color: var(--edna-orange);
  font-weight: bold;
}

.btn-close {
  position: absolute;
  top: 5px;
  right: 5px;
  background: none;
  color: var(--edna-red);
  font-size: 1.2rem;
  cursor: pointer;
}
</style>

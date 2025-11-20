// src/services/api.js
import axios from 'axios';

// Cria uma instância do axios que já aponta para o seu back-end
const apiClient = axios.create({
    baseURL: 'http://localhost:8080/api/v1', // O BasePath que definimos!
    headers: {
        'Content-Type': 'application/json',
    },
});

export default {
    // TODO: Resto das funções de acesso à API 
    getFornecedores(filters = null) {
        return apiClient.get('/fornecedores');
    },
    getFornecedorById(id) {
        return apiClient.get(`/fornecedores/${id}`);
    },
    createFornecedor(data) {
        // data é um objeto JS, ex: { nome: "...", cnpj: "..." }
        return apiClient.post('/fornecedores', data);
    },
    deleteFornecedor(id) {
        return apiClient.delete(`/fornecedores/${id}`);
    },
    getProdutos(filters = null) {
        // filters pode ser um objeto, ex: { params: { 'filter-nome': 'ilike.Cerveja' } }
        return apiClient.get('/produtos', filters);
    },
    getClientes(filters = null) {
        return apiClient.get('/clientes');
    },
    getOfertas(filters = null) {
        return apiClient.get('/ofertas')
    }
};

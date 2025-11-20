<script setup>
</script>

<template>
  <div id="layout">
    <header>
      <nav class="navbar">
        <ul class="nav-links">
          <li><RouterLink to="/">Início</RouterLink></li>
          <li><RouterLink to="/clientes">Clientes</RouterLink></li>
          <li><RouterLink to="/produtos">Produtos</RouterLink></li>
          <li><RouterLink to="/ofertas">Ofertas</RouterLink></li>
          <li><RouterLink to="/fornecedores">Fornecedores</RouterLink></li>
        </ul>
      </nav>
    </header>

    <main class="content">
      <RouterView />
    </main>
  </div>
</template>

<style>
.navbar {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  background-color: transparent;
  color: white;
  position: absolute;
  width: 100vw;
  border: 0px;
  margin: 1rem 0rem;
}
.nav-brand {
  font-weight: bold;
}
.nav-links {
  list-style: none;
  display: flex;
  padding: 0px 2rem;
}
.nav-links li {
  padding: 15px 20px;
}
.nav-links a {
  color: #0d0d0d;
  text-decoration: none;
  font-weight: 500;
}
.nav-links a:hover,
.nav-links a.router-link-active {
  color: #F5E8BA;
}
</style>

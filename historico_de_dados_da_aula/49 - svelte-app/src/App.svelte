<script>
	import { Router, Route } from 'svelte-routing';
	import Login from './pages/login/Index.svelte'
	import Menu from './componentes/Menu.svelte'
	import Donos from './pages/donos/Index.svelte'
	import NovoDono from './pages/donos/Salvar.svelte'
	import AlterarDono from './pages/donos/Salvar.svelte'
	import Pets from './pages/pets/Index.svelte';
	import Administradores from './pages/administradores/Index.svelte';
	import Home from './pages/home/Index.svelte';
	import AuthGuard from './guard/AuthGuard.svelte';
</script>

<main>
	<Router>
		<Route path="login" component={Login} />
		
		<Route path="">
			<AuthGuard>
				<Menu />
				<Home />
			</AuthGuard>
		</Route>
		<Route path="donos">
			<AuthGuard>
				<Menu />
				<Donos />
			</AuthGuard>
		</Route>
		<Route path="donos/novo">
			<AuthGuard>
				<Menu />
				<NovoDono />
			</AuthGuard>
		</Route>
		<Route path="donos/:id" let:params>
			<AuthGuard>
				<Menu />
				<AlterarDono id="{params.id}" />
			</AuthGuard>
		</Route>
		<Route path="pets">
			<AuthGuard>
				<Menu />
				<Pets />
			</AuthGuard>
		</Route>
		<Route path="administradores">
			<AuthGuard>
				<Menu />
				<Administradores />
			</AuthGuard>
		</Route>
	</Router>
</main>

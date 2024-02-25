<script>
	import { Router, Route, Link } from 'svelte-routing';
	import Donos from './pages/donos/Index.svelte'
	import Pets from './pages/pets/Index.svelte';
	import Administradores from './pages/administradores/Index.svelte';
	import Home from './pages/home/Index.svelte';

	let currentPath = "/";
    function clickMenu(item) {
        currentPath = item;
    }
</script>

<main>
	<Router>
		<header data-bs-theme="dark">
			<nav class="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
			  <div class="container-fluid">
				<Link class="navbar-brand" to="" on:click={() => clickMenu('/')}>Svelte Golang</Link>
				<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
				  <span class="navbar-toggler-icon"></span>
				</button>
				<div class="collapse navbar-collapse" id="navbarCollapse">
				  <ul class="navbar-nav me-auto mb-2 mb-md-0">

					<li class="nav-item"><Link class="nav-link {currentPath === '/' ? 'active' : ''}" on:click={() => clickMenu('/')} to="">Home</Link></li>
					<li class="nav-item"><Link class="nav-link {currentPath === '/donos' ? 'active' : ''}" on:click={() => clickMenu('/donos')} to="donos">Donos</Link></li>
					<li class="nav-item"><Link class="nav-link {currentPath === '/pets' ? 'active' : ''}" on:click={() => clickMenu('/pets')} to="pets">Pets</Link></li>
					<li class="nav-item"><Link class="nav-link {currentPath === '/administradores' ? 'active' : ''}" on:click={() => clickMenu('/administradores')} to="administradores">Administradores</Link></li>
				  </ul>
				</div>
			  </div>
			</nav>
		</header>
		
		<Route path="" component={Home} />
		<Route path="donos" component={Donos} />
		<Route path="pets" component={Pets} />
		<Route path="administradores" component={Administradores} />
	</Router>
</main>

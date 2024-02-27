<script>
	import { Link, navigate } from 'svelte-routing';

	let currentPath = window.location.pathname;
    const clickMenu = (item) => {
        currentPath = item;
    }

	const sair = () => {
		localStorage.clear();
		navigate("/login");
	}
</script>

<header data-bs-theme="dark">
    <nav class="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
      <div class="container-fluid">
        <Link class="navbar-brand" to="" on:click={() => clickMenu('/')}>Svelte Golang</Link>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarCollapse">
          <ul class="navbar-nav me-auto mb-2 mb-md-0">

            <li class="nav-item"><Link class={`nav-link ${currentPath === '/' ? 'active' : ''}`} on:click={() => clickMenu('/')} to="">Home</Link></li>
            <li class="nav-item"><Link class={`nav-link ${currentPath === '/donos' ? 'active' : ''}`} on:click={() => clickMenu('/donos')} to="donos">Donos</Link></li>
            <li class="nav-item"><Link class={`nav-link ${currentPath === '/pets' ? 'active' : ''}`} on:click={() => clickMenu('/pets')} to="pets">Pets</Link></li>
            <li class="nav-item"><Link class={`nav-link ${currentPath === '/administradores' ? 'active' : ''}`} on:click={() => clickMenu('/administradores')} to="administradores">Administradores</Link></li>
            <li class="nav-item"><Link class={`nav-link ${currentPath === '/recursos' ? 'active' : ''}`} on:click={() => clickMenu('/recursos')} to="recursos">Recursos</Link></li>
            <li class="nav-item"><button class="nav-link" on:click={sair}>Sair</button></li>
          </ul>
        </div>
      </div>
    </nav>
</header>
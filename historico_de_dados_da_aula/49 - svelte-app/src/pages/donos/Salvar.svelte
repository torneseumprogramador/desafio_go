<script>
    import { API_HOST } from '../../config';
    import { navigate, Link } from 'svelte-routing'

	export let id;

	if(id){
		(async function() {
			try {
				const token = localStorage.getItem('userToken');

				const resposta = await fetch(`${API_HOST}/donos/${id}`, {
					method: 'GET',
					headers: {
						'Authorization': `Bearer ${token}`,
						'Content-Type': 'application/json'
					}
				});

				if (!resposta.ok) {
					throw new Error(`Erro na requisição: ${resposta.status}`);
				}

				const dono = await resposta.json();
				nome = dono.Nome
				telefone = dono.Telefone
			} catch (erro) {
				console.error("Falha ao carregar os dono:", erro);
				throw erro;
			}
		})();

	}

    let nome = '';
    let telefone = '';

	/*  === exemplo de como carregar dados em um select box ===
    let donoId;

	let donos = []; // Lista de donos

    async function carregarDonos() {
        try {
			const token = localStorage.getItem('userToken');

			const resposta = await fetch(`${API_HOST}/donos`, {
				method: 'GET',
				headers: {
					'Authorization': `Bearer ${token}`,
					'Content-Type': 'application/json'
				}
			});

			if (!resposta.ok) {
				throw new Error(`Erro na requisição: ${resposta.status}`);
			}

			donos = await resposta.json();
		} catch (erro) {
			console.error("Falha ao carregar os donos:", erro);
			throw erro;
		}
    }

	(async function() {
		carregarDonos()
	})();
	*/

    let mensagemErro = '';

    async function salvarDono(event) {
        event.preventDefault();

        try {
			const token = localStorage.getItem('userToken');

			let url = `${API_HOST}/donos`;
			let method = 'POST';
			if(id){
				url = `${API_HOST}/donos/${id}`;
				method = 'PUT';
			}

			const resposta = await fetch(url, {
				method: method,
				headers: {
					'Authorization': `Bearer ${token}`,
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ Nome: nome, Telefone: telefone })
			});

            const resultado = await resposta.json();

            if (!resposta.ok) {
				mensagemErro = resultado.erro;
				return;
            }

            navigate('/donos');
        } catch (erro) {
            mensagemErro = erro.message;
        }
    }
</script>

<main>
    <h1>Adicionar Novo Dono</h1>
    <form on:submit={salvarDono}>
        <div class="form-group">
            <label for="nome">Nome:</label>
            <input type="text" class="form-control" id="nome" bind:value={nome}>
        </div>
        <div class="form-group">
            <label for="telefone">Telefone:</label>
            <input type="text" class="form-control" id="telefone" bind:value={telefone}>
        </div>

        <!--div class="form-group">
            <label for="pets">Pets:</label>
            <select class="form-control" id="donoId" bind:value={donoId}>
				<option value="">[Selecione]</option>
				{#each donos as dono}
					<option value="{dono.Id}">{dono.Nome}</option>
				{/each}
			</select>
        </div-->

		{#if mensagemErro}
			<div class="alert alert-danger mt-3" role="alert">
				{mensagemErro}
			</div>
		{/if}

		<br>

        <button type="submit" class="btn btn-primary">Salvar</button>
		<Link to="/donos" class="btn btn-default">Voltar</Link>
    </form>
</main>

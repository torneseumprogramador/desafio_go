<script>
    import { API_HOST } from '../../config';
    import { navigate, Link } from 'svelte-routing'

	export let id;

	if(id){
		(async function() {
			try {
				const token = localStorage.getItem('userToken');

				const resposta = await fetch(`${API_HOST}/recursos/${id}`, {
					method: 'GET',
					headers: {
						'Authorization': `Bearer ${token}`,
						'Content-Type': 'application/json'
					}
				});

				if (!resposta.ok) {
					throw new Error(`Erro na requisição: ${resposta.status}`);
				}

				const recurso = await resposta.json();
				nome = recurso.nome
				valor = recurso.valor
			} catch (erro) {
				console.error("Falha ao carregar os recurso:", erro);
				throw erro;
			}
		})();

	}

    let nome = '';
    let valor = '';

    let mensagemErro = '';

    async function salvarRecurso(event) {
        event.preventDefault();

        try {
			const token = localStorage.getItem('userToken');

			let url = `${API_HOST}/recursos`;
			let method = 'POST';
			if(id){
				url = `${API_HOST}/recursos/${id}`;
				method = 'PUT';
			}

			const resposta = await fetch(url, {
				method: method,
				headers: {
					'Authorization': `Bearer ${token}`,
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ Nome: nome, Valor: valor })
			});

            const resultado = await resposta.json();

            if (!resposta.ok) {
				mensagemErro = resultado.erro;
				return;
            }

            navigate('/recursos');
        } catch (erro) {
            mensagemErro = erro.message;
        }
    }
</script>

<main>
    <h1>Adicionar Novo Recurso</h1>
    <form on:submit={salvarRecurso}>
        <div class="form-group">
            <label for="nome">Nome:</label>
            <input type="text" class="form-control" id="nome" bind:value={nome}>
        </div>
        <div class="form-group">
            <label for="valor">Valor:</label>
            <input type="number" class="form-control" id="valor" bind:value={valor}>
        </div>

		{#if mensagemErro}
			<div class="alert alert-danger mt-3" role="alert">
				{mensagemErro}
			</div>
		{/if}

		<br>

        <button type="submit" class="btn btn-primary">Salvar</button>
		<Link to="/recursos" class="btn btn-default">Voltar</Link>
    </form>
</main>

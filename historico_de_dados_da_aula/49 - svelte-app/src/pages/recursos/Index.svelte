<script>
    import { API_HOST } from '../../config';
	import { Link } from 'svelte-routing';

	$: recursos = [];
	
	const carregarRecursos = async () => {
		try {
			const token = localStorage.getItem('userToken');

			const resposta = await fetch(`${API_HOST}/recursos`, {
				method: 'GET',
				headers: {
					'Authorization': `Bearer ${token}`,
					'Content-Type': 'application/json'
				}
			});

			if (!resposta.ok) {
				throw new Error(`Erro na requisição: ${resposta.status}`);
			}

			recursos = await resposta.json();
		} catch (erro) {
			console.error("Falha ao carregar os recursos:", erro);
			throw erro;
		}
	};

	const excluirItem = async (id) => {
		if(confirm("Confirma ?")){
			try {
				const token = localStorage.getItem('userToken');

				const resposta = await fetch(`${API_HOST}/recursos/${id}`, {
					method: 'DELETE',
					headers: {
						'Authorization': `Bearer ${token}`,
						'Content-Type': 'application/json'
					}
				});

				if (!resposta.ok) {
					throw new Error(`Erro na requisição: ${resposta.status}`);
				}

				carregarRecursos();
			} catch (erro) {
				console.error("Falha ao carregar os recursos:", erro);
				throw erro;
			}
		}
	};


	(async function() {
		carregarRecursos()
	})();

</script>

<main>
    <h1>Recursos da api Golang com Fiber</h1>
	<hr>
	<Link to="/recursos/novo" class="btn btn-primary">Adicionar</Link>
	<hr>
    <div class="table-responsive">
        <table class="table table-striped table-hover">
            <thead class="table-dark">
                <tr>
                    <th>Nome</th>
                    <th>Telefone</th>
                    <th colspan="2">Acões</th>
                </tr>
            </thead>
            <tbody>
                {#each recursos as recurso}
                <tr>
                    <td>{recurso.nome}</td>
                    <td>{recurso.valor}</td>
                    <td style="width: 50px">
						<Link to="/recursos/{recurso.id}" class="btn btn-warning">Alterar</Link>
					</td>
					<td style="width: 50px">
						<button on:click={()=> { excluirItem(recurso.id) }} class="btn btn-danger">Excluir</button>
					</td>
                </tr>
                {/each}
            </tbody>
        </table>
    </div>
</main>

<script>
    import { API_HOST } from '../../config';
	import { Link } from 'svelte-routing';

	$: donos = [];
	
	const carregarDonos = async () => {
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
	};

	const excluirItem = async (id) => {
		if(confirm("Confirma ?")){
			try {
				const token = localStorage.getItem('userToken');

				const resposta = await fetch(`${API_HOST}/donos/${id}`, {
					method: 'DELETE',
					headers: {
						'Authorization': `Bearer ${token}`,
						'Content-Type': 'application/json'
					}
				});

				if (!resposta.ok) {
					throw new Error(`Erro na requisição: ${resposta.status}`);
				}

				carregarDonos();
			} catch (erro) {
				console.error("Falha ao carregar os donos:", erro);
				throw erro;
			}
		}
	};


	(async function() {
		carregarDonos()
	})();

</script>

<main>
    <h1>Donos</h1>
	<hr>
	<Link to="/donos/novo" class="btn btn-primary">Adicionar</Link>
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
                {#each donos as dono}
                <tr>
                    <td>{dono.Nome}</td>
                    <td>{dono.Telefone}</td>
                    <td style="width: 50px">
						<Link to="/donos/{dono.Id}" class="btn btn-warning">Alterar</Link>
					</td>
					<td style="width: 50px">
						<button on:click={()=> { excluirItem(dono.Id) }} class="btn btn-danger">Excluir</button>
					</td>
                </tr>
                {/each}
            </tbody>
        </table>
    </div>
</main>

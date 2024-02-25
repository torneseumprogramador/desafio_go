<script>
	$: donos = [];
    import { API_HOST } from '../../config';
	
	const carregarDonos = async () => {
		try {
			const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRhbmlsby5hcGFyZWNpZG8uc2FudG9zQGdtYWlsLmNvbSIsImV4cCI6MTcwODg1ODQxMiwiaWQiOiJhMzhlZDgwMC04NjIyLTQwY2YtODdmZC03NmU0OTYwNDYxYWUiLCJub21lIjoiRGFuaWxvIiwic3VwZXIiOnRydWV9.0xcLBGlL-UKLtaXZrcCahy28-wnsXsf4xECbOZIRwEs';

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

			const donos = await resposta.json();

			return donos;
		} catch (erro) {
			console.error("Falha ao carregar os donos:", erro);
			throw erro;
		}
	};


	(async function() {
		donos = await carregarDonos()
	})();

</script>

<main>
    <h1>Donos</h1>
	<hr>
    <div class="table-responsive">
        <table class="table table-striped table-hover">
            <thead class="table-dark">
                <tr>
                    <th>Nome</th>
                    <th>Telefone</th>
                </tr>
            </thead>
            <tbody>
                {#each donos as dono}
                <tr>
                    <td>{dono.Nome}</td>
                    <td>{dono.Telefone}</td>
                </tr>
                {/each}
            </tbody>
        </table>
    </div>
</main>

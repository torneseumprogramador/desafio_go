<script>
    import { navigate } from 'svelte-routing'
    import { API_HOST } from '../../config';

    let email = '';
    let senha = '';

    let mensagemErro = '';

    async function Logar(event) {
        event.preventDefault();

        try {
			const resposta = await fetch(`${API_HOST}/login`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ Email: email, Senha: senha })
			});

            const resultado = await resposta.json();

            if (!resposta.ok) {
				mensagemErro = resultado.erro;
				return;
            }

			localStorage.setItem('userToken', resultado.token);
            navigate('/');
        } catch (erro) {
            mensagemErro = erro.message;
        }
    }
</script>

<main>
    <h1>Login</h1>
    <form on:submit={Logar}>
        <div class="form-group">
            <label for="email">Email:</label>
            <input type="text" class="form-control" id="email" bind:value={email}>
        </div>
        <div class="form-group">
            <label for="senha">Senha:</label>
            <input type="password" class="form-control" id="senha" bind:value={senha}>
        </div>

		{#if mensagemErro}
			<div class="alert alert-danger mt-3" role="alert">
				{mensagemErro}
			</div>
		{/if}

		<br>

        <button type="submit" class="btn btn-primary">Logar</button>
    </form>
</main>

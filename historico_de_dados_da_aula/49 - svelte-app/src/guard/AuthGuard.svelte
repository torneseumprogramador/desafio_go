<!-- src/AuthGuard.svelte -->
<script>
    import { onMount } from 'svelte';
    import { navigate } from 'svelte-routing';
    import { isAuthenticated } from './guards';
    
    onMount(async () => {
        const autenticado = await isAuthenticated();
        if (!autenticado) {
            navigate('/login', { replace: true });
        }
    });
</script>

<slot></slot>

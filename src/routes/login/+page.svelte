<script lang="ts">
    import { onMount } from 'svelte';
    import { login } from "$lib/auth";

    let email: HTMLInputElement;
    let password: HTMLInputElement;

    let loginButtonEnabled = false;

    onMount(() => {
        email = document.getElementById('form-email') as HTMLInputElement;
        password = document.getElementById('form-password') as HTMLInputElement;
    });

    const onClickLoginButton = async () => {
        await login(email.value.trim(), password.value.trim());

        email.value = "";
        password.value = "";
    };

    const onInput = () => {
        toggleEnabledLoginButton();
    }

    const toggleEnabledLoginButton = () => {
        loginButtonEnabled = email.value.trim().length > 0 && password.value.trim().length > 0;
    }
</script>

<h1>login</h1>
<div>
    <input type="text" id="form-email" placeholder="e-mail" on:input={onInput} />
    <input type="password" id="form-password" placeholder="password" on:input={onInput} />
    <button type="button" on:click={onClickLoginButton} disabled="{!loginButtonEnabled}">login</button>
</div>
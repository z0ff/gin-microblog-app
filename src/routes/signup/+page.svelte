<script lang="ts">
    import { onMount } from 'svelte';
    import { signup, login } from "$lib/auth";

    let name: HTMLInputElement;
    let displayname: HTMLInputElement;
    let email: HTMLInputElement;
    let password: HTMLInputElement;

    let loginButtonEnabled = false;

    onMount(() => {
        name = document.getElementById('form-name') as HTMLInputElement;
        displayname = document.getElementById('form-displayname') as HTMLInputElement;
        email = document.getElementById('form-email') as HTMLInputElement;
        password = document.getElementById('form-password') as HTMLInputElement;
    });

    const onClickSignupButton = async () => {
        await signup(name.value.trim(), displayname.value.trim(), email.value.trim(), password.value.trim());
        await login(email.value.trim(), password.value.trim());

        name.value = "";
        displayname.value = "";
        email.value = "";
        password.value = "";
    };

    const onInput = () => {
        toggleEnabledLoginButton();
    }

    const toggleEnabledLoginButton = () => {
        loginButtonEnabled = name.value.trim().length > 0 && displayname.value.trim().length > 0 && email.value.trim().length > 0 && password.value.trim().length > 0;
    }
</script>

<h1>Signup</h1>
<div>

    <input type="text" id="form-name" placeholder="name" on:input={onInput} />
    <input type="text" id="form-displayname" placeholder="displayname" on:input={onInput} />
    <input type="text" id="form-email" placeholder="e-mail" on:input={onInput} />
    <input type="password" id="form-password" placeholder="password" on:input={onInput} />
    <button type="button" on:click={onClickSignupButton} disabled="{!loginButtonEnabled}">signup</button>
</div>

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
        const result = await login(email.value.trim(), password.value.trim());

        name.value = "";
        displayname.value = "";
        email.value = "";
        password.value = "";

        if (result === false) {
            alert("login failed");
        } else {
            location.href = "/";
        }
    };

    const onInput = () => {
        toggleEnabledLoginButton();
    }

    const toggleEnabledLoginButton = () => {
        loginButtonEnabled = name.value.trim().length > 0 && displayname.value.trim().length > 0 && email.value.trim().length > 0 && password.value.trim().length > 0;
    }
</script>

<div class="w-full flex justify-center items-center">
    <div class="card card-body flex flex-column max-w-sm">
        <input type="text" id="form-name" placeholder="name" on:input={onInput} class="input" />
        <input type="text" id="form-displayname" placeholder="displayname" on:input={onInput} class="input" />
        <input type="text" id="form-email" placeholder="e-mail" on:input={onInput} class="input" />
        <input type="password" id="form-password" placeholder="password" on:input={onInput} class="input" />
        <button type="button" on:click={onClickSignupButton} disabled="{!loginButtonEnabled}" class="btn btn-primary">会員登録</button>
        <a class="btn btn-primary" href="/login">ログイン</a>
    </div>
</div>

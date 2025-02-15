<script>
    import { strToDate } from "$lib/datetime";
    export let post;
</script>
<div class="card card-bordered bg-base-50 shadow-lg">
                <div class="card-body">
                    <p class="font-bold">{post.User.DisplayName}</p>
                    <p>{strToDate(post.CreatedAt)}</p>
                    <p class="whitespace-pre-wrap">{post.Content}</p>
                </div>
            </div>

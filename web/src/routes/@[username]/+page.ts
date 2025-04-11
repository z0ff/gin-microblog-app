import type { PageLoad } from '@sveltejs/kit';

export const load: PageLoad = async ({ params }) => {
    return {
        props: {
            username: params.username
        }
    }
}

export default async ({ store, router }) => {
    store.dispatch("auth/myProfile")
        .catch((err) => {
            router.push("/login");
        });

}
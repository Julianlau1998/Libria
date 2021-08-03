<template>
    <div>
        <div>
            <img :src="$auth.user.picture">
            <h2>{{ $auth.user.name }}</h2>
            <p>{{ $auth.user.email }}</p>
        </div>

            <pre>
                {{ JSON.stringify($auth.user, null, 2) }}
            </pre>
    </div>
</template>

<script>
export default {

}
</script>

<style>

</style>
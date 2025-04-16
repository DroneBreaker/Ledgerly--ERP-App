<template>
    <div class="justify-self-center">
        <h1 class="text-5xl">
            {{ loginMsg }}
        </h1>

        <form action="#">
            <input type="text" placeholder="Email or Username"/>
        </form>
    </div>
</template>

<script setup>
/**
 * IMPORTS
 */


/**
 * PROPS
 */
 const props = defineProps({
    loginMsg: {
        type: String,
        default: "Login to Ledgerly",
    }
 })


</script>

<style lang="scss" scoped>

</style>
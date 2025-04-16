<template>
    <button class="text-sm p-2 px-4 bg-white text-black rounded-md w-28 h-12">
        {{ buttonText }}
    </button>
</template>

<script setup>
/**
 * IMPORTS
*/


/**
 * PROPS
*/
 const props = defineProps({
    buttonText: {
        type: String,
        default: "Button",
        required: true
    }
 })
</script>

<style lang="scss" scoped>

</style>
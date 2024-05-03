<script setup>
import { UserCircleIcon } from "@heroicons/vue/24/solid";
</script>

<template>
    <nav class="bg-gray-800 text-sky-600 flex items-center justify-between px-4 border-b-[1px] border-b-gray-900 sticky top-0 z-50">
        <h3>
            <strong>Snap</strong>
            <i>cart</i>
        </h3>
        <UserCircleIcon class="text-slate-50 h-8 w-8" />
    </nav>
</template>

<style scoped>
    nav{
        height: 64px;
    }
</style>
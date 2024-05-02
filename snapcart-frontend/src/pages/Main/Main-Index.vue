<script setup>
    import ProductCard from '@/components/product/Product-Card.vue';
import { ref } from 'vue';

    const data = [
        {
            title : "Satu"
        },
        {
            title : "Dua"
        },
        {
            title : "Tiga"
        },
        {
            title : "Empat"
        },
        {
            title : "Lima"
        },
        {
            title : "Enam"
        },
        {
            title : "Tujuh"
        },
        {
            title : "Dua"
        },
    ]

    const onDetail = ref(false)
    const toggleDetail = () => onDetail.value = !onDetail.value
</script>

<template>
    <main class="container mx-auto lg:px-12">
        <section class="flex gap-4 overflow-auto">
            <div class="grid-cols-4 gap-4 py-4 grid w-full" :class="{'w-3/4' : !onDetail, 'grid-cols-3' : onDetail}">
                <ProductCard v-for="(val, idx) in data" :key="idx" :title="val.title" :toggle="toggleDetail" />
            </div>
            <div class="w-1/3 bg-red-600 transition-all text-white overflow-hidden" :class="{'w-0' : !onDetail}">
                <span>testing</span>
            </div>  
        </section>
    </main>
</template>
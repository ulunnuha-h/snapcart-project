<script setup>
    import ProductCard from '@/components/product/Product-Card.vue';
    import { ref } from 'vue';
    import { StarIcon } from "@heroicons/vue/24/solid";

    const sideDetail = ref(false);
    const activeId = ref(0);
    const toggleSideDetail = (id) => {
        if(activeId.value == id || !sideDetail.value){
            sideDetail.value = !sideDetail.value
        }
        activeId.value = id
        console.log(id);
    } 

    const data = [
        {
            title : "Sepatu kulit merk Sony dan dapat dilipat seperti messi ",
            price : 30000,
            rating : 4.5,
            img : "https://picsum.photos/id/30/1280/901"
        },
        {
            title : "Robot Megatron yang sangat canggih",
            price : 75000,
            rating : 4.1,
            img : "https://picsum.photos/id/32/4032/3024"
        },
        {
            title : "Celana pendek",
            price : 40000,
            rating : 3.5,
            img : "https://picsum.photos/id/36/4179/2790"
        },
        {
            title : "Yihaa apa ini",
            price : 33000,
            rating : 4.2,
            img : "https://picsum.photos/id/42/3456/2304"
        },
    ]
</script>

<template>
    <main class="min-h-screen">
        <section 
            :class="{'gap-4':sideDetail}"
            class="flex overflow-auto h-full ">
            <div 
                :class="{'w-2/3' : sideDetail}" 
                class="grid grid-cols-4 gap-4 py-4 h-full w-full justify-around">
                <ProductCard v-for="(val, idx) in data" :key="idx" :data="val" :toggle="toggleSideDetail" :id="idx"/>
            </div>
            <div 
                :class="{'w-1/3 h-fit opacity-100 border-l-[1px] border-l-gray-900 p-4':sideDetail}" 
                class="opacity-0 overflow-hidden transition-all  w-0 h-0">
                <h3 class="text-gray-100 font-bold rounded-md">{{data[activeId].title}}</h3>
                <img class="w-full aspect-square object-cover my-4 rounded-md" :src="data[activeId].img" alt="">
                <section class="flex justify-between items-center">
                    <h2 class="text-gray-100 font-semibold">{{"Rp " +data[activeId].price.toLocaleString() }}</h2>
                    <span class="flex items-center gap-1 text-amber-400 text-base">
                        <StarIcon class="h-4 w-4" />
                        <span>{{data[activeId].rating}}/5</span>
                    </span>
                </section>
                <hr class="border-gray-900 my-4 ">
                <p class="text-gray-200 text-justify mb-4">Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin vel commodo velit. Aenean mauris lorem, volutpat maximus nisi quis, lobortis pharetra risus. Suspendisse interdum id ipsum volutpat elementum. Praesent sed sem sit amet ante ultrices laoreet sed quis quam. Donec porttitor libero nec nibh cursus volutpat. Nullam luctus dui arcu, eu consequat dui interdum sit amet. Nam enim lorem, aliquam in mollis sed, eleifend ac ipsum. Duis ut bibendum dolor. Quisque eu nisl vestibulum, dignissim arcu a, congue mauris. Proin nibh quam, elementum auctor lobortis nec, viverra tempus enim. In a tristique nisl. Integer ut nisi porttitor, bibendum dolor id, suscipit turpis. Nullam tincidunt quam ullamcorper luctus ullamcorper. Aliquam ligula magna, hendrerit id laoreet at, accumsan efficitur orci.</p>
                <section class="flex gap-4">
                    <button class="btn-primary py-2 px-4 w-full">Add to cart</button>
                    <button class="btn-primary py-2 px-4 w-full">Add to cart</button>
                </section>
            </div>
        </section>
    </main>
</template>
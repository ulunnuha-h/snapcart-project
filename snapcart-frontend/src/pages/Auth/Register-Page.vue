<script setup>
import { ref } from 'vue';
import { EyeIcon, EyeSlashIcon } from "@heroicons/vue/24/solid";
const data = ref([
  {
    label : "Email address",
    id : "email",
    type : "email",
    toggle : false,
    value : ""
  },
  {
    label : "Full name",
    id : "fullname",
    type : "text",
    toggle : false,
    value : ""
  },
  {
    label : "Password",
    id : "password",
    type : "password",
    toggle : true,
    value : ""
  },
  {
    label : "Confirm password",
    id : "confirmPassword",
    type : "password",
    toggle : true,
    value : ""
  },
])

const toggleHandler = (idx) => {
  data.value[idx].type = (data.value[idx].type  == "password" ? "text" : "password")
}

</script>
<template>
    <div class="flex min-h-screen flex-col justify-center px-6 py-12 lg:px-8 bg-gray-900">
      <div class="sm:mx-auto sm:w-full sm:max-w-sm">
        <h2 class="mt-10 text-center font-bold leading-9 tracking-tight text-gray-200">Register your account</h2>
      </div>
    
      <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
        <form class="space-y-6">
          <div v-for="(val, idx) in data" :key="idx">
            <label class="block text-sm font-medium leading-6 text-gray-200">{{ val.label }}</label>
            <div class="mt-2 relative flex items-center">
              <input 
                :value="data[idx].value" 
                @input="e => data[idx].value = e.target.value"
                :id="val.id" 
                :name="val.id" 
                :type="val.type" 
                required 
                class="block w-full rounded-md border-0 py-1.5 px-2.5 placeholder:text-gray-400 sm:text-sm sm:leading-6 bg-gray-800 ring-1 ring-gray-700 text-white">
                <span v-if="val.toggle" class="absolute right-2 h-6 w-6 text-white cursor-pointer">
                  <EyeSlashIcon v-if="val.type=='password'" @click="()=>toggleHandler(idx)"/>
                  <EyeIcon v-else @click="()=>toggleHandler(idx)"/>
                </span>
            </div>
          </div>
          <div>
            <button type="submit" class="flex w-full btn-primary">Sign up</button>
          </div>
        </form>
    
        <p class="mt-10 text-center text-sm text-gray-500">
            Already a member?
          <RouterLink to="/login" class="font-semibold leading-6 text-sky-600 hover:text-sky-500">Login Now</RouterLink>
        </p>
      </div>
    </div>
</template>
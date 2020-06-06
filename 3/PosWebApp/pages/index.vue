<template>
  <div>
    <b-navbar toggleable type="dark" variant="dark">
      <b-navbar-brand href="#">
        NavBar
      </b-navbar-brand>

      <b-button variant="warning">
        ADD
      </b-button>
    </b-navbar>

    <b-container fluid>
        <b-table striped hover :items="items" />
    </b-container>
    
  </div>
</template>
<script>
export default {
  data () {
    return {
      items: [
        { age: 40, first_name: 'Dickerson', last_name: 'Macdonald' },
        { age: 21, first_name: 'Larsen', last_name: 'Shaw' },
        { age: 89, first_name: 'Geneva', last_name: 'Wilson' },
        { age: 38, first_name: 'Jami', last_name: 'Carney' }
      ]
    }
  }
}
</script>

<template>
  <div>
    <b-navbar toggleable type="dark" variant="dark">
      <b-navbar-brand href="#">
        NavBar
      </b-navbar-brand>

      <b-button v-b-modal.modal-1 variant="warning">
        ADD
      </b-button>
    </b-navbar>

    <b-container fluid>
        <b-table striped hover :items="items" />
    </b-container>

		<b-modal id="modal-1" title="Add">
			<b-form @submit="onSubmit" @reset="onReset">
        <b-form-group
          id="input-group-1"
          label="Nama Kasir"
          label-for="input-1"
        >
          <b-form-input
            id="input-1"
            type="email"
            required
            placeholder="Masukkan Nama Kasir"
          />
        </b-form-group>

				<b-form-group
          id="input-group-1"
          label="Jenis Pesanan"
          label-for="input-1"
        >
          <b-form-input
            id="input-1"
            type="email"
            required
            placeholder="Masukkan Jenis Pesanan"
          />
        </b-form-group>

				<b-form-group
          id="input-group-1"
          label="Nama Product"
          label-for="input-1"
        >
          <b-form-input
            id="input-1"
            type="email"
            required
            placeholder="Masukkan Nama Product"
          />
        </b-form-group>

				<b-form-group
          id="input-group-1"
          label="Harga"
          label-for="input-1"
        >
          <b-form-input
            id="input-1"
            type="email"
            required
            placeholder="Masukkan Harga"
          />
        </b-form-group>

      <!-- <b-form-group id="input-group-2" label="Your Name:" label-for="input-2">
        <b-form-input
          id="input-2"
          v-model="form.name"
          required
          placeholder="Enter name"
        ></b-form-input>
      </b-form-group>

      <b-form-group id="input-group-3" label="Food:" label-for="input-3">
        <b-form-select
          id="input-3"
          v-model="form.food"
          :options="foods"
          required
        ></b-form-select>
      </b-form-group> -->

      <!-- <b-form-group id="input-group-4">
        <b-form-checkbox-group v-model="form.checked" id="checkboxes-4">
          <b-form-checkbox value="me">Check me out</b-form-checkbox>
          <b-form-checkbox value="that">Check that out</b-form-checkbox>
        </b-form-checkbox-group>
      </b-form-group> -->

        <b-button type="submit" variant="primary">Submit</b-button>
    </b-form>
		</b-modal>
    
  </div>
</template>
<script>
export default {
  data () {
    return {
      items: [
        { no: 1, cashier: 'Pevite Pearce', product: 'Latte', category: 'Drink', price: "Rp. 10.000" },
        { no: 2, cashier: 'Pevite Pearce', product: 'Latte', category: 'Drink', price: "Rp. 10.000" },
        { no: 3, cashier: 'Pevite Pearce', product: 'Latte', category: 'Drink', price: "Rp. 10.000" },
        { no: 4, cashier: 'Pevite Pearce', product: 'Latte', category: 'Drink', price: "Rp. 10.000" },
      ]
    }
  }
}
</script>

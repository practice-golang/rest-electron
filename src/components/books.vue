<template>
  <div>
    <div id="book-data">
      <input type="hidden" name="id" ref="id" placeholder="0">
      <span>Title / Author:</span>
      <input name="Title" ref="Title_new" placeholder>
      <input name="Author" ref="Author_new" placeholder>
      <button @click="createUpdateItem" id="btn-insert-update">Insert/Change Data</button>
      <p></p>

      <button @click="getList(null)" id="call_book_data">Get book list</button>
      
      <img :class="isLoading ? 'show' : 'hide'" src="../assets/loading.gif">
      
      <span v-html="data_loading"></span>
      <table border="1">
        <tr>
          <th colspan="6">Book List</th>
        </tr>
        <tr v-for="(items, i) in book_data" :key="i">
          <template v-if="edit === items.ID">
            <td>
              <input v-model="items.ID" type="disabled" name="id" ref="id">
            </td>
            <td>
              <input v-model="items.Title" name="book_title" ref="book_title">
            </td>
            <td>
              <input v-model="items.Author" name="author" ref="author">
            </td>
            <td>
              <button @click="createUpdateItem(items)" id="update_book_data">Apply</button>
            </td>
            <td>
              <button @click="cancelEdit()" id="cancel_edit">Cancel</button>
            </td>
          </template>
          <template v-else>
            <td>{{items.ID}}</td>
            <td>{{items.Title}}</td>
            <td>{{items.Author}}</td>
            <td>
              <button @click="editItem(items)" id="edit_book_data">Edit</button>
            </td>
            <td>
              <button @click="deleteItem(items.ID)" id="delete_book_data">Delete</button>
            </td>
          </template>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: "book-data",
  data: function () {
    return {
      isLoading: false,
      rest_url: "http://localhost:1323",
      data_loading: "",
      book_data: [],
      edit: 0
    }
  },
  mounted() {
    this.isLoading = false
  },
  methods: {
    getList: function () {
      var vm = this
      var xhr = new XMLHttpRequest()

      this.isLoading = true

      xhr.onreadystatechange = function () {
        xhr.readyState > 3 && xhr.status == 200
          ? (vm.book_data = JSON.parse(xhr.responseText))
          : null
      }

      this.isLoading = false

      xhr.open("GET", this.rest_url + "/books", true)
      xhr.send()
    },
    createUpdateItem: function (item) {
      var vm = this
      var xhr = new XMLHttpRequest()

      var mode = ""
      var params = ""

      if (item.ID != undefined) {
        vm.$data.edit = 0;
        params = "ID=" + item.ID + "&Title=" + item.Title + "&Author=" + item.Author;
        mode = "PUT";
        // console.log(params);
      } else {
        params =
          "Title=" +
          vm.$refs.Title_new.value +
          "&Author=" +
          vm.$refs.Author_new.value;
        mode = "POST";
      }

      this.isLoading = true

      xhr.onreadystatechange = function () {
        xhr.readyState > 3 && xhr.status == 200
          ? ((vm.data_loading = ""), alert(xhr.responseText), vm.getList())
          : null
      }

      this.isLoading = false;

      xhr.open(mode, this.rest_url + "/books", true);
      xhr.setRequestHeader(
        "Content-Type",
        "application/x-www-form-urlencoded; charset=UTF-8"
      )
      xhr.send(params);
    },
    editItem: function (book) {
      this.$data.edit = book.ID;
    },
    cancelEdit: function () {
      this.$data.edit = 0
    },
    deleteItem: function (book_id) {
      var vm = this
      var xhr = new XMLHttpRequest();

      // this.data_loading = "<img src='../assets/loading.gif' height=21>";
      xhr.onreadystatechange = function () {
        xhr.readyState > 3 && xhr.status == 200
          ? ((vm.data_loading = ""), alert(xhr.responseText), vm.getList())
          : null
      }

      xhr.open("DELETE", this.rest_url + "/books/" + book_id, true)
      xhr.send()
    }
  }
}
</script>

<style>
.show {
  display: block;
}
.hide {
  display: none;
}
</style>
<template>
  <q-header class="bg-white text-grey-10">
    <q-toolbar class="constrain x">

      <!-- Home Button -->
      <q-btn flat to="/">
        <q-icon left size="3em" name="eva-camera-outline" />
        <q-toolbar-title class="text grand-hotel text-bold">
          Home
        </q-toolbar-title>
      </q-btn>

      <q-separator class="large-screen-only" vertical spaced />

      <!-- Search -->
      <q-toolbar-title class="text-center">
        <q-input
          v-model="searchText"
          bottom-slots
          class="nuks"
          label="Search"
          @keyup.enter="GoSearch"
        />
      </q-toolbar-title>

      <!-- Chat Button -->
      <q-btn
        round
        @click="GoToChat"
        :icon="unReadedMessages > 0 ? 'eva-message-square-outline' : 'eva-message-square'"
        :color="unReadedMessages > 0 ? 'primary' : 'dark'"
      >
        <q-badge
          v-if="unReadedMessages > 0"
          color="negative"
          floating
          rounded
          :label="unReadedMessages"
        />
      </q-btn>

      <!-- Notification Button -->
      <q-btn
        round
        @click="GoToNotification"
        :icon="notificationNum > 0 ? 'eva-bell-outline' : 'eva-bell'"
        :color="notificationNum > 0 ? 'primary' : 'dark'"
      >
        <q-badge
          v-if="notificationNum > 0"
          floating
          color="negative"
          rounded
          :label="notificationNum"
        />
      </q-btn>

      <!-- User Avatar -->
      <q-btn round>
        <q-avatar size="42px">
          <img src="https://cdn-icons-png.flaticon.com/512/3237/3237472.png" />
        </q-avatar>

        <q-menu>
          <q-list style="min-width: 100px">
            <q-item clickable v-close-popup>
              <q-item-section @click="Profile">
                Profile
              </q-item-section>
            </q-item>

            <q-separator />

            <q-item clickable v-close-popup>
              <q-item-section @click="LogUserOut">
                Logout
              </q-item-section>
            </q-item>
          </q-list>
        </q-menu>
      </q-btn>

    </q-toolbar>
  </q-header>
</template>

<script>
export default {
  name: 'NavBar',

  data() {
    return {
      searchText: '',
      notificationNum: 2,      // Static demo number
      unReadedMessages: 3      // Static demo number
    }
  },

  methods: {
    GoSearch() {
      console.log("Search:", this.searchText)
      this.$router.push({
        path: `/Search`,
        query: { search: this.searchText }
      })
    },

    Profile() {
      this.$router.push(`/Profile/demo-user`)
    },

    LogUserOut() {
      console.log("Logged out (frontend only)")
      this.$router.push(`/Auth`)
    },

    GoToNotification() {
      this.$router.push('/Notification')
    },

    GoToChat() {
      this.$router.push('/Chat')
    }
  }
}
</script>

<style lang="sass">
.nuks 
  width: 250px
  text-align: center
  display: inline-block !important

.q-toolbar-title
  display: flex 
  align-items: center 

.q-btn 
  margin-left: 10px
</style>

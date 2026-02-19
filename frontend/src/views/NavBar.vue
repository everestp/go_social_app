<template>
  <q-header class="bg-white text-grey-9" bordered>
    <q-toolbar class="constrain">

      <!-- Logo / Home -->
      <q-btn flat to="/" class="q-px-md">
        <q-icon name="eva-home-outline" size="28px" class="q-mr-sm" />
        <q-toolbar-title class="text-bold text-h6">
          Home
        </q-toolbar-title>
      </q-btn>

      <q-space />

      <!-- Search -->
      <q-input
        dense
        outlined
        rounded
        placeholder="Search..."
        class="search-input"
        @keyup.enter="GoSearch"
      >
        <template v-slot:prepend>
          <q-icon name="eva-search-outline" />
        </template>
      </q-input>

      <q-space />

      <!-- Messages -->
      <q-btn round flat>
        <q-icon name="eva-message-circle-outline" size="24px" />
        <q-badge
          v-if="unReadedMessages > 0"
          color="negative"
          floating
          rounded
          :label="unReadedMessages"
        />
      </q-btn>

      <!-- Notifications -->
      <q-btn round flat>
        <q-icon name="eva-bell-outline" size="24px" />
        <q-badge
          v-if="notificationNum > 0"
          color="negative"
          floating
          rounded
          :label="notificationNum"
        />
      </q-btn>

      <!-- Profile -->
      <q-btn round flat>
        <q-avatar size="36px">
          <img src="https://cdn-icons-png.flaticon.com/512/3237/3237472.png" />
        </q-avatar>

        <q-menu>
          <q-list style="min-width: 150px">
            <q-item clickable v-close-popup>
              <q-item-section avatar>
                <q-icon name="eva-person-outline" />
              </q-item-section>
              <q-item-section>Profile</q-item-section>
            </q-item>

            <q-separator />

            <q-item clickable v-close-popup>
              <q-item-section avatar>
                <q-icon name="eva-log-out-outline" />
              </q-item-section>
              <q-item-section>Logout</q-item-section>
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
  props: {
    unReadedMessages: {
      type: Number,
      default: 0
    },
    notificationNum: {
      type: Number,
      default: 0
    }
  },
  methods: {
    GoSearch () {
      console.log('Search triggered')
    }
  }
}
</script>

<style lang="sass">
.search-input
  width: 280px

.q-btn
  margin-left: 8px
</style>

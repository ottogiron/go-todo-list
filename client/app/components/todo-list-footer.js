import Ember from 'ember';

export default Ember.Component.extend({
  actions: {
    clearCompleted() {
      var completed = this.get('completed');
      completed.invoke('deleteRecord');
      completed.invoke('save');
    }
  }
});

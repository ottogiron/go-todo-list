import Ember from 'ember';

export default Ember.Component.extend({
  canToggle: function () {
    var anyTodos = this.get('todos.length');
    var isEditing = this.get('todos').isAny('isEditing');
    return anyTodos && !isEditing;
  }.property('todos.length', '@each.isEditing')
});

import Ember from 'ember';

export default Ember.Controller.extend({
		todos: Ember.inject.controller(),
		allTodos: Ember.computed.alias('todos'),
		itemController: 'todo',
		canToggle: function () {
			var anyTodos = this.get('allTodos.length');
			var isEditing = this.isAny('isEditing');
			return anyTodos && !isEditing;
		}.property('allTodos.length', '@each.isEditing')
	});

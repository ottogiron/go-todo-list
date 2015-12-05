import Ember from 'ember';

export default Ember.Component.extend({
    store: Ember.inject.service(),
		actions: {
			createTodo() {
				var title, todo;

				// Get the todo title set by the "New Todo" text field
				title = this.get('newTitle').trim();
				if (!title) {
					return;
				}

				// Create the new Todo model
				todo = this.get('store').createRecord('todo', {
					title: title,
					isCompleted: false
				});
				todo.save();

				// Clear the "New Todo" text field
				this.set('newTitle', '');
			},

			clearCompleted() {
				var completed = this.get('completed');
				completed.invoke('deleteRecord');
				completed.invoke('save');
			}
		},

		/* properties */

		remaining: Ember.computed.filterBy('model', 'isCompleted', false),
		completed: Ember.computed.filterBy('model', 'isCompleted', true),

		allAreDone: Ember.computed('length', 'completed.length', {
			get: () => {
				var length = this.get('length');
				var completedLength = this.get('completed.length');
				return length > 0 && length === completedLength;
			},
			set: (key, value) => {
				this.setEach('isCompleted', value);
				return value;
			}
		})
	});


# workflows

This document is a tentative outline of general workflows.

_In the future it will be broken down by service, and the lists will be replaced with diagrams._


## login form workflows

- click forgot password
	- if email is empty display error
	- if email is populated display sending
		- if service does not find email display error
		- if service finds email display sent

- supplied valid email
	- green checkbox is displayed
	- login button is enabled

- supplied invalid email
	- green plus sign is displayed
	- form displays hidden fields:
		- password confirmation
		- display-name
	- register button is enabled

_Thus far it appears valid passwords are at least 8 characters long, contain one capital, one number, and one special character._

_The email validation is very basic, contains **only** one `@`, at least one `.`, and is at least 5 characters long._

_Additional server-side validation will be used, but this prevents silly mistakes and bad form submissions._


- oauth link is clicked
	- approval must be provided for redirect
	- system grabs token and acquires identity
	- system looks for existing identity
		- if found, login is complete, localStorage is cleared and updated data is polled
		- if not found, page displays login form as registration form and auto-populates with acquired values
			- hidden fields are used to associate with oauth

Individual services are responsible for a set of DOM elements; for example the credentials system is responsible for managing the display of the login form, and any pre-populated fields or ajax callbacks.  It is also responsible for clearing localStorage whenever a login or logout event occurs, to prevent stale or invalid data.  A menu system will be responsible for updating the menu items according to the state of the credentials (eg. are they are user, are they an administrator, is the administrator status active?).

(function(window, document) {
	// primary js for csideanime

    // class-manipulation helper functions
    var addClass = function(o, c) {
        var classes = o.getAttribute('class') || '';
        if (classes.indexOf(c) != -1) return;
        o.setAttribute('class', classes + ' ' + c);
    };
    var removeClass = function(o, c) {
        var classes = o.getAttribute('class') || '';
        classes = classes.replace(c, '').trim();
        o.setAttribute('class', classes);
    };
    var hasClass = function(o, c) {
        var classes = o.getAttribute('class') || '';
        return classes.indexOf(c) != -1 ?  true : false;
    }
    var likeAnEmail= function(s) {
        var e = s.indexOf('@');
        return /[.]+/.test(s) && s.length > 4 && (e != 0 && e != -1 && e != s.length-1);
    }
    var containsNumbers = function(s) {
        return /\d/.test(s);
    };
    var containsCapitals = function(s) {
        return /[A-Z]/.test(s);
    };
    var containsSpecialCharacters = function(s) {
        return /[^A-Za-z0-9]+/.test(s);
    };

    // check for login form
    var loginFormGroup = document.getElementsByClassName('loginForm');
    if (loginFormGroup.length > 0) {
        var loginForm = loginFormGroup[0];

        // grab email input, login buttons, and forgot-password link
        var signIn = loginForm.signIn;
        var register = loginForm.register;
        var email = loginForm.email;
        var forgotPassword = loginForm.getElementsByTagName('a')[0];

        // define reusable functions
        var removeMessages = function() {
            var messages = loginForm.getElementsByClassName('messages');
            for (var i = 0; i < messages.length; i++) {
                messages[i].parentNode.removeChild(messages[i]);
            }
        }
        var disableButtons = function(e, t) {
            addClass(signIn, 'disabled');
            addClass(register, 'disabled');
            email.parentNode.setAttribute('class', '');
            removeClass(loginForm, 'register');
            removeMessages();
        };
        var validateIdentity = function() {
            disableButtons();

            // ignore short input
            if (email.value.length <= 3) return;

            // @todo(casey): ajax call to grab emails

            // @test-code
            validEmails = [];
            if (email.value.length > 3) {
                validEmails.push('cdelorme@gmail.com');
            }

            // @todo(casey): move this into a callback that validates against real returns
            if (email.value.length > 3) {
                if (validEmails.indexOf(email.value) != -1) {
                    removeClass(signIn, 'disabled');
                    addClass(email.parentNode, 'after-ok');
                } else if (likeAnEmail(email.value)) {
                    removeClass(register, 'disabled');
                    addClass(email.parentNode, 'after-add');
                    addClass(loginForm, 'register');
                }
            }
        };
        var handleSubmit = function(e) {
            e.preventDefault();
            console.log("form validation...");
            removeMessages();

            var validationMessages = [];
            if (!likeAnEmail(loginForm.email.value)) {
                validationMessages.push('invalid email');
            }
            if (loginForm.email.length <= 3) {
                validationMessages.push('please supply an email...');
            }
            if (loginForm.password.value.length < 8) {
                validationMessages.push('password is less than 8 characters...');
            }
            if (!containsCapitals(loginForm.password.value)) {
                validationMessages.push('password has no capital letters...');
            }
            if (!containsNumbers(loginForm.password.value)) {
                validationMessages.push('password has no numbers...');
            }
            if (!containsSpecialCharacters(loginForm.password.value)) {
                validationMessages.push('password has no special characters...');
            }
            if (!hasClass(register, 'disabled')) {
                // verify email
                if (loginForm.passwordConfirmation.value != loginForm.password.value) {
                    validationMessages.push('password confirmation does not match the supplied password...');
                }
                if (loginForm.displayName.value.length <= 3) {
                    validationMessages.push('display name is too short...');
                }
                if (validationMessages.length == 0) {
                    var message = document.createElement('div');
                    addClass(message, 'messages success');
                    message.appendChild(document.createTextNode('sending registration request...'));
                    loginForm.insertBefore(message, loginForm.firstChild);
                    // @todo(casey): ajax class to register
                }
            } else if (!hasClass(signIn, 'disabled') && validationMessages.length == 0) {
                var message = document.createElement('div');
                addClass(message, 'messages success');
                message.appendChild(document.createTextNode('sending login request...'));
                loginForm.insertBefore(message, loginForm.firstChild);
                // @todo(casey): ajax class to login
            }
            if (validationMessages.length > 0) {
                var message = document.createElement('div');
                addClass(message, 'messages error');
                for (var i = 0; i < validationMessages.length; i++) {
                    var pm = document.createElement('p');
                    pm.appendChild(document.createTextNode(validationMessages[i]));
                    message.appendChild(pm);
                }
                loginForm.insertBefore(message, loginForm.firstChild);
            }
        };
        var handleForgottenPassword = function(e) {
            e.preventDefault();
            removeMessages();

            // verify email is populated
            if (email.value.length <= 3) {
                var message = document.createElement('div');
                message.appendChild(document.createTextNode('please enter a username'));
                addClass(message, 'messages error');
                addClass(email.parentNode, 'after-error');
                loginForm.insertBefore(message, loginForm.firstChild);
            } else {
                var message = document.createElement('div');
                message.appendChild(document.createTextNode('processing reset...'));
                addClass(message, 'messages info');
                loginForm.insertBefore(message, loginForm.firstChild);

                // send ajax call
                // update notice on callback to check email
            }
        }

        // handle forgot password link
        forgotPassword.addEventListener('click', handleForgottenPassword, false);

        // watch email input
        email.addEventListener('blur', validateIdentity, false);
        email.addEventListener('keyup', validateIdentity, false);

        // handle form submission
        signIn.addEventListener('click', handleSubmit, false);
        register.addEventListener('click', handleSubmit, false);
        loginForm.addEventListener('submit', handleSubmit, false);

        // logging
        console.log('login form activated');
    }

})(window, document);

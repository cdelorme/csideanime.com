
# pages

First Release Pages/Views (Mobile-First Member Perspective):

- Login /w google/facebook integration
- Registration
- Member Profile
- Settings for
    - Profile
    - Theme/Style
- Home /w Threads by Tags
- View Thread /w posts
- Create Thread

Using immediate-render mode, we can store user permissions as part of the app state, and render control elements on-demand once the objects are on-screen.  This avoids the click-or-tap-to-display concerns with mobile systems.  We should do what we can to keep the UI uncluttered.

Members with special permissions who we may want to be able to see things like ip addresses and ban buttons should have a toggle at the top to turn their special controls on.  This let's them have the standard uncluttered members experience by default, which will probably be cleaner and preferrable.


// JavaScript, um das Calendly-Widget beim Klicken des Buttons anzuzeigen
document.getElementById('showCalendly').addEventListener('click', function () {
document.getElementById('calendlyWidgetContainer').style.display = 'block';
document.getElementById('calendlyWidget').style.display = 'block';
this.style.display = 'none'; // Versteckt den Button nach dem Klicken
});
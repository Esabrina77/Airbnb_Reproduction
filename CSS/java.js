function playMusic(audioId) {
    var audio = document.getElementById(audioId);
    audio.play();
  }

  function convertToLowerCase(input) {
    input.value = input.value.toLowerCase();
}
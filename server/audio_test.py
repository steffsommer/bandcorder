import sounddevice as sd
import soundfile as sf

def get_default_mic():
  mic_index = sd.default.device[0]
  default_mic = sd.query_devices()[mic_index]
  print(default_mic)
  return default_mic
  
get_default_mic()
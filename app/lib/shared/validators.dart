String? validateIPv4(String? value) {
  if (value == null || value.isEmpty) {
    return 'IPv4 address is required';
  }

  final segments = value.split('.');
  if (segments.length != 4) {
    return 'Invalid IPv4 address format';
  }

  for (final segment in segments) {
    final number = int.tryParse(segment);
    if (number == null || number < 0 || number > 255) {
      return 'Each IPv4 segment must be between 0 and 255';
    }

    if (segment.length > 1 && segment.startsWith('0')) {
      return 'Leading zeros are not allowed';
    }
  }

  return null;
}

import 'package:bandcorder/models/recording_status.dart';
import 'package:flutter/material.dart';

class RecordingStateInfo extends StatelessWidget {
  final RecordingState? _recordingState;

  RecordingStateInfo({
    super.key,
    required RecordingState? recordingState,
  }) : _recordingState = recordingState;

  @override
  Widget build(BuildContext context) {
    if (_recordingState != null) {
      return GridView.count(
        crossAxisCount: 1,
        shrinkWrap: true,
        childAspectRatio: 4.0,
        children: [
          _buildInfoTile('💾 File name', _recordingState.fileName),
          _buildInfoTile('🕒 Duration', _recordingState.duration.toString()),
        ],
      );
    } else {
      return const Text('');
    }
  }

  Widget _buildInfoTile(String label, String value) {
    return Card(
      elevation: 2.0,
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              label,
              style: const TextStyle(
                fontSize: 14.0,
                fontWeight: FontWeight.bold,
                color: Colors.grey,
              ),
            ),
            const SizedBox(height: 4.0),
            Text(
              value,
              style: const TextStyle(
                fontSize: 20.0,
                fontWeight: FontWeight.w500,
              ),
              overflow: TextOverflow.ellipsis,
            ),
          ],
        ),
      ),
    );
  }
}

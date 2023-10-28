import React, { useState } from 'react';
import {QrScanner} from '@yudiel/react-qr-scanner';
import { Typography, Container, Paper, Button, Divider } from '@mui/material';

const Dashboard = () => {
  const [attendanceRecord, setAttendanceRecord] = useState(null);
  const [scanning, setScanning] = useState(false);

  const handleScan = (data) => {
    console.log('aaaaa',data);
    
    if (data) {
      setAttendanceRecord(data);
      setScanning(false);
    }
  };

  const handleError = (err) => {
    console.error(err);
  };

  const startScanning = () => {
    setScanning(true);
  };

  return (
    <Container maxWidth="sm">
      <Paper elevation={3} style={{ padding: '16px', textAlign: 'center' }}>
        <Typography variant="h4" gutterBottom>
          صفحه حضور و عیاب
        </Typography>
        <Divider style={{ margin: '16px 0' }} />
        {attendanceRecord ? (
          <div>
            <Typography variant="body1" gutterBottom>
              ثبت حضور یا عیاب با موفقیت انجام شد.
            </Typography>
            <Typography variant="body2" gutterBottom>
              زمان: {new Date().toLocaleString()}
            </Typography>
            <Typography variant="body2" gutterBottom>
              وضعیت: {attendanceRecord}
            </Typography>
          </div>
        ) : (
          <div>
            {scanning ? (
         
              <QrScanner
              onError={handleError}
              onScan={handleScan}
              onDecode={e=>console.log('bbbbb',e)
              }
      />
            ) : (
              <div>
                <Typography variant="body1" gutterBottom>
                  لطفاً QR Code را اسکن کنید.
                </Typography>
                <Button
                  variant="contained"
                  color="primary"
                  onClick={startScanning}
                  style={{ marginTop: '16px' }}
                >
                  شروع اسکن
                </Button>
              </div>
            )}
          </div>
        )}
      </Paper>
    </Container>
  );
};

Dashboard.getLayout = function getLayout(page: ReactElement) {
  return <>{page}</>;
};

export default Dashboard;
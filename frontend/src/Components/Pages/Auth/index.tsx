import { Button, TextField, Typography, Container, Box, Paper } from '@mui/material';
import { useForm, Controller } from 'react-hook-form';

const Auth = () => {
  const { control, handleSubmit } = useForm();
  
  const onSubmit = (data) => {
    console.log(data);
  };

  return (
    <Container maxWidth="sm">
      <Paper elevation={3} style={{ padding: '16px', textAlign: 'center' }}>


      <form onSubmit={handleSubmit(onSubmit)}>
        <Typography variant="h4" gutterBottom>
          Login
        </Typography>
        <Box component="div" m={2}>
          <Controller
            name="email"
            control={control}
            defaultValue=""
            render={({ field }) => (
              <TextField
                {...field}
                label="Email"
                variant="outlined"
                fullWidth
              />
            )}
          />
        </Box>
        <Box component="div" m={2}>
          <Controller
            name="password"
            control={control}
            defaultValue=""
            render={({ field }) => (
              <TextField
                {...field}
                type="password"
                label="Password"
                variant="outlined"
                fullWidth
              />
            )}
          />
        </Box>
        <Box component="div" m={2}>
          <Button type="submit" variant="contained" color="primary">
            Login
          </Button>
        </Box>
      </form>


      </Paper>

    </Container>
  );
};

Auth.getLayout = function getLayout(page) {
  return <>{page}</>;
};

export default Auth;

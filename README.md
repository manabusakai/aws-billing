# aws-billing

## Configuration

Please set either.
(See also: [Configuring the AWS Command Line Interface](http://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html))

### Environment Variables

```
$ export AWS_ACCESS_KEY_ID=YOUR_ACCESS_KEY_ID
$ export AWS_SECRET_ACCESS_KEY=YOUR_SECRET_ACCESS_KEY
```

### Credential Files

```
[default]
aws_access_key_id = YOUR_ACCESS_KEY_ID
aws_secret_access_key = YOUR_SECRET_ACCESS_KEY
```

## Usage

```
$ aws-billing
Billing: 18.83
```

## Enable Billing Alerts

This command needs to enable billing alerts.
So, turn on "[Receive Billing Alerts](https://console.aws.amazon.com/billing/home?#/preferences)".

## License

This software is released under the MIT License.

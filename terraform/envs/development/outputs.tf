output "instance_public_ip" {
    value = {
        api     = aws_instance.ec2-api.public_ip
        client  = aws_instance.ec2-client.public_ip
        storage = aws_instance.ec2-storage.public_ip
        indexer = aws_instance.ec2-indexer.public_ip
    }
}
